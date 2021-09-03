package config

import (
	"encoding/json"
	"os"
)

// Configuration stores setting values
type Configuration struct {
	port string `json:"port"`
}

// Config shares the global configuration
var (
	Config *Configuration
)

// LoadConfig loads configuration from the config file
func LoadConfig() error {
	// Filename is the path to the json config file
	file, err := os.Open("front/config/config.json")
	if err != nil {
		return err
	}

	Config = new(Configuration)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		return err
	}

	return nil
}

func (c *Configuration) Port() string {
	if c.port != "" {
		return c.port
	}
	return ":8080"
}
