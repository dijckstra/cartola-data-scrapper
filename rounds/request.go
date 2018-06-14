package rounds

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dijckstra/cartola-data-scrapper/data/database"
	"github.com/dijckstra/cartola-data-scrapper/data/model"

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
	Db *database.DB
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
	var res []model.Round
	err = json.Unmarshal(body, &res)
	if err != nil {
		panic(err)
	}

	// get finished rounds
	getFinishedRounds(&res)
	err = requestor.Db.InsertRounds(res)
	if err != nil {
		panic(err)
	}

	// request each finished round
	requestor.requestMatchesPerformed(len(res))
}

// getFinishedRounds returns the slice of rounds that were played until now.
func getFinishedRounds(rounds *[]model.Round) {
	var count int

	for _, r := range *rounds {
		if time.Until(r.EndTime.Time) > 0 {
			break
		}

		count++
	}

	*rounds = (*rounds)[:count]
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

		_, err2 := core.DecodeJSON(body)
		if err != nil {
			panic(err2)
		}

		responses++
		if responses == count {
			break
		}
	}
}
