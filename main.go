package main

import (
	"github.com/dijckstra/cartola-data-scrapper/request"
	"github.com/jasonlvhit/gocron"
)

func main() {
	gocron.Every(5).Seconds().Do(request.CurrentRound)
	<-gocron.Start()
}
