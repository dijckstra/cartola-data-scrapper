package main

import (
	"log"

	"github.com/jasonlvhit/gocron"

	"github.com/dijckstra/cartola-data-scrapper/request"
)

func main() {
	log.Printf("startou")
	// request player information every Monday
	playerRequestor := &request.PlayerRequestor{}
	
	playerRequestor.RequestPlayers()
	gocron.Every(1).Monday().Do(playerRequestor.RequestPlayers)
	<-gocron.Start()
}
