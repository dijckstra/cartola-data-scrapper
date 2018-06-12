package main

import (
	"github.com/dijckstra/cartola-data-scrapper/rounds"
	"github.com/jasonlvhit/gocron"
)

func main() {
	roundRequestor := rounds.NewRoundRequestor()

	gocron.Every(5).Seconds().Do(roundRequestor.RequestMatchesPerformed)
	<-gocron.Start()
}
