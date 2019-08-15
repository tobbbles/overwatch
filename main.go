package main

import (
	"fmt"
	"go.uber.org/zap"
	"service/environment"
	"service/remote/overwatch"
	"service/server"
	"service/worker"
	"time"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Sprintf("failed initialising zap logger: %s", err))
	}
	defer logger.Sync()

	// Environment variable defaulting and setting
	env, err := environment.Load()
	if err != nil {
		logger.Fatal("failed loading environment", zap.Error(err))
	}

	// Create a remote overwatch client
	owc, err := overwatch.New()
	if err != nil {
		logger.Fatal("failed creating overwatch api client", zap.Error(err))
	}

	// Setup and start our remote api worker
	controller, err := worker.New(owc, 5*time.Second)
	if err != nil {
		logger.Fatal("failed creating overwatch api client worker controller", zap.Error(err))
	}

	// Start our controller in it's go routine since it blocks until error
	go func() {
		if err := controller.Start(); err != nil {
			panic(err)
		}
	}()

	// Configure and setup the HTTP server with it's endpoints.
	config := &server.Config{
		Addr:   env.Address,
		Logger: logger,
	}

	s, err := server.New(config)

	if err != nil {
		logger.Fatal("failed creating server", zap.Error(err))
	}

	logger.Info("starting server", zap.String("address", config.Addr))

	if err := s.Start(); err != nil {
		logger.Fatal("fatal server error", zap.Error(err))
	}
}
