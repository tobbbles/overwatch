package worker

import (
	"fmt"
	"log"
	"service/remote/overwatch"
	"time"
)

type Controller struct {
	ticker *time.Ticker
	client *overwatch.Client

	quit chan error
}

func New(client *overwatch.Client, interval time.Duration) (*Controller, error) {

	controller := &Controller{
		ticker: time.NewTicker(interval),
		client: client,
	}

	return controller, nil
}

func (c *Controller) start() {
	for range c.ticker.C {

		count, err := c.client.HeroCount()
		if err != nil {
			c.quit <- err
			return
		}


		log.Printf("count: %d\n", count)

		// Iterate through the hero IDs
		for i := 1; i < count; i++ {
			hero, err := c.client.Hero(i)
			if err != nil {
				c.quit <- err
				return
			}

			// TODO: Insert hero and it's abilities
			fmt.Printf("%+v\n", hero)
		}
	}
}

// Start will start the worker controller and block until it returns.
func (c *Controller) Start() error {

	log.Println("starting controller")

	go c.start()

	select {
	case err := <-c.quit:
		return err
	}
}
