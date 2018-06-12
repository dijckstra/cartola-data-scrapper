package rounds

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

// RoundRequestor represents the object that executes requests to the
// rounds and matches endpoints.
type RoundRequestor struct {
	adapter RoundAdapter
}

// NewRoundRequestor returns an instance of a RoundRequestor.
func NewRoundRequestor() RoundRequestor {
	roundRequestor := RoundRequestor{}
	roundRequestor.adapter = RoundAdapter{}

	return roundRequestor
}

// RequestMatchesPerformed executes requests to every match played until now.
func (requestor *RoundRequestor) RequestMatchesPerformed() {
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

	// get finished rounds
	rounds := getFinishedRounds(json)

	// request each finished round
	requestor.requestMatchesPerformed(len(rounds))
}

// getFinishedRounds returns the slice of rounds that were played until now.
func getFinishedRounds(j []core.JSON) []core.JSON {
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

	return j[:count]
}

func (requestor *RoundRequestor) requestMatchesPerformed(count int) {
	ch := make(chan *http.Response)

	for i := 1; i <= count; i++ {
		go func(round int) {
			url := fmt.Sprintf(roundEnpoint, round)
			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}

			ch <- resp
		}(i)
	}

	requestor.responseHandler(ch, count)
}

func (requestor *RoundRequestor) responseHandler(ch <-chan *http.Response, count int) {
	responses := 0

	for {
		resp := <-ch

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		json, err := core.DecodeJSON(body)
		if err != nil {
			panic(err)
		}

		requestor.adapter.GetRound(json)

		responses++
		if responses == count {
			break
		}
	}
}
