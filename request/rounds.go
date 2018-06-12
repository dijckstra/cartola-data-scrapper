package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dijckstra/cartola-data-scrapper/core"
)

const (
	roundsEndpoint = "https://api.cartolafc.globo.com/rodadas"
	roundEnpoint   = "https://api.cartolafc.globo.com/partidas/%d"
	timeFormat     = "2006-01-02 15:04:05"
)

// MatchesPerformed executes requests to every match played until now.
func MatchesPerformed() {
	// request rounds list
	resp, err := http.Get(roundsEndpoint)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// parse to JSON
	json, err := core.DecodeJSONArray(body)
	if err != nil {
		panic(err)
	}

	// get number of finished rounds
	count := getFinishedRounds(json)

	// request each finished round
	requestMatchesPerformed(count)
}

// getFinishedRounds returns the number of rounds that
// were played until now.
func getFinishedRounds(j []core.JSON) int {
	var roundTimeString *string
	var roundTime time.Time
	var count int

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

func requestMatchesPerformed(count int) {
	ch := make(chan *http.Response)

	for i := 1; i <= count; i++ {
		go func(count int) {
			url := fmt.Sprintf(roundEnpoint, count)
			resp, err := http.Get(url)
			ch <- resp

			defer resp.Body.Close()

			if err != nil {
				panic(err)
			}
		}(count)
	}
}
