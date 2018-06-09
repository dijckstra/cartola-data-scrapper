package request

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dijckstra/cartola-data-scrapper/core"
)

const (
	roundsEndpoint = "https://api.cartolafc.globo.com/rodadas"
	roundEnpoint   = "https://api.cartolafc.globo.com/partidas/%s"
	timeFormat     = "2006-01-02 15:04:05"
)

// CurrentRound executes a request to the currently active round.
func CurrentRound() {
	resp, err := http.Get(roundsEndpoint)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	json, err := core.DecodeJSONArray(body)
	if err != nil {
		panic(err)
	}

	// count := GetFinishedRounds(json)
}

// GetFinishedRounds returns the number of rounds that
// were played until now.
func GetFinishedRounds(j []core.JSON) int32 {
	var roundTimeString *string
	var roundTime time.Time
	var count int32

	for _, r := range j {
		roundTimeString = core.StringFromJSON(r, "fim")
		roundTime, _ = time.Parse(timeFormat, *roundTimeString)

		if time.Until(roundTime) > 0 {
			break
		}

		count++
	}

	return count
}
