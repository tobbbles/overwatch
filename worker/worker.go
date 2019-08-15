package worker

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"service/remote/overwatch"
	"service/store"
	"time"
)

type Controller struct {
	ticker  *time.Ticker
	client  *overwatch.Client
	updater store.Updater
	logger  *zap.Logger

	errors chan error
	quit   chan struct{}
}

// Config used to instantiate the Controller
type Config struct {
	Client   *overwatch.Client
	Interval time.Duration
	Logger   *zap.Logger
	Updater  store.Updater
}

// New takes an Overwatch client and an Updater - it'll update the store with data from the overwatch client
// at an interval specified.
func New(config *Config) (*Controller, error) {
	if config.Client == nil {
		return nil, fmt.Errorf("%T.Client can not be nil when creating a new controller", config)
	}
	if config.Updater == nil {
		return nil, fmt.Errorf("%T.Updater can not be nil when creating a new controller", config)
	}
	if config.Interval == 0 {
		return nil, fmt.Errorf("%T.Interval can not be zero when creating a new controller", config)
	}

	controller := &Controller{
		ticker:  time.NewTicker(config.Interval),
		client:  config.Client,
		updater: config.Updater,
		logger:  config.Logger,

		quit:   make(chan struct{}),
		errors: make(chan error, 8),
	}

	return controller, nil
}

func (c *Controller) start() {
	// Intial tick
	c.tick()
	c.logger.Info("complete inital fetch of remote overwatch api data")

	// Begin ticker
	for range c.ticker.C {
		c.logger.Info("begining refresh of remote overwatch api data")
		c.tick()
	}
}

// Tick is called on every iteration of the ticker
func (c *Controller) tick() {
	count, err := c.client.HeroCount()
	if err != nil {
		c.errors <- err
		return
	}

	// Iterate through the hero IDs
	for i := 1; i < count; i++ {
		hero, err := c.client.Hero(i)
		if err != nil {
			c.errors <- err
			return
		}

		// Update the store via the updater
		if err := c.updater.Update(hero); err != nil {
			c.errors <- err
			continue
		}
	}
}

// Start will start the worker controller and block until it returns.
func (c *Controller) Start() error {

	log.Println("starting controller")

	go c.start()

	for {
		select {
		case err := <-c.errors:
			c.logger.Error("error from worker controller", zap.Error(err))

		case <-c.quit:
			return nil
		}
	}
}
