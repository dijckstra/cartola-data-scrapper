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
	if _, err := toml.DecodeFile("/go/src/github.com/dijckstra/cartola-data-scrapper/config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
