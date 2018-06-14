package main

import (
	"log"

	"github.com/dijckstra/cartola-data-scrapper/config"
	"github.com/dijckstra/cartola-data-scrapper/data"
	"github.com/dijckstra/cartola-data-scrapper/rounds"
)

var configuration = config.Configuration{}

// Parse the configuration file 'config.toml'
func init() {
	configuration.Read()
}

func main() {
	// Establish a connection to DB
	db, err := data.NewDB(configuration.Server, configuration.Database)
	if err != nil {
		log.Panic(err)
	}

	roundRequestor := &rounds.RoundRequestor{Db: db}
	roundRequestor.RequestMatchesPerformed()
}
