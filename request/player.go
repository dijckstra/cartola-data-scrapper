package request

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	playersEndpoint = "https://api.cartolafc.globo.com/atletas/mercado"
	contentType     = "application/json"
)

// PlayerRequestor is the object that
// retrieves player information from the API.
type PlayerRequestor struct{}

// RequestPlayers retrieves all players in the tournament.
func (requestor *PlayerRequestor) RequestPlayers() {
	// request rounds list
	resp, err := http.Get("http://naming-service:4000/lookup?name=cartola-rest-api")
	if err != nil {
		panic(err)
	}
	// read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	selfEndpoint := string(body) + "/players"
	
	resp, err = http.Get(playersEndpoint)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// read response
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// parse to JSON
	var objmap map[string]*json.RawMessage
	err = json.Unmarshal(body, &objmap)
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(*objmap["atletas"])
	if err != nil {
		panic(err)
	}

	// send to our endpoint
	resp, err = http.Post(selfEndpoint, contentType, bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	log.Printf("sucesso")
}
