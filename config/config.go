package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Configuration represents the TOML configuration file.
type Configuration struct {
	Server   string
	Database string
}

// Read will parse the configuration file.
func (c *Configuration) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
