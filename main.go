package main

import (
	"log"

	"github.com/jasonlvhit/gocron"

	"github.com/dijckstra/cartola-data-scrapper/config"
	"github.com/dijckstra/cartola-data-scrapper/data/database"
	"github.com/dijckstra/cartola-data-scrapper/request"
)

var configuration = config.Configuration{}

// Parse the configuration file 'config.toml'
func init() {
	configuration.Read()
}

func main() {
	// Establish a connection to DB
	db, err := database.NewDB(configuration.Server, configuration.Database)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("startou")
	// request player information every Monday
	playerRequestor := &request.PlayerRequestor{Db: db}

	playerRequestor.RequestPlayers()
	gocron.Every(1).Monday().Do(playerRequestor.RequestPlayers)
	<-gocron.Start()
}
