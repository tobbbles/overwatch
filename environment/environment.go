package environment

import (
	"github.com/codingconcepts/env"
)

// Config contains all configurable variables for tweaking and configuring this service
// from environment variables
type Config struct {
	// Address is an optional configuration to change the bind address
	Address string `env:"ADDR" default:":3000"`

	// SQLitePath for the store
	SQLitePath string `env:"SQLITE_PATH" default:"./overwatch.db"`

	// Interval in minutes we want to fetch fresh data from Overwatch API
	Interval int `env:"INTERVAL" default:"30"`
}

// Load environment variables into the config.
func Load() (*Config, error) {
	c := &Config{}

	if err := env.Set(c); err != nil {
		return nil, err
	}

	return c, nil
}
