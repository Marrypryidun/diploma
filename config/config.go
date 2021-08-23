package config

import (
	"encoding/json"
	"os"
)

// Configuration stores setting values
type Configuration struct {
	port    string `json:"port"`
	mgAddrs string `json:"mgAddrs"`
	dbName  string `json:"dbName"`
}

// Config shares the global configuration
var (
	Config *Configuration
)

// LoadConfig loads configuration from the config file
func LoadConfig() error {
	// Filename is the path to the json config file
	file, err := os.Open("config/config.json")
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
func (c *Configuration) MgAddrs() string {
	if c.mgAddrs != "" {
		return c.mgAddrs
	}
	return "mongodb://127.0.0.1"
}

func (c *Configuration) MgDbName() string {
	if c.mgAddrs != "" {
		return c.dbName
	}
	return "Diplom"
}
