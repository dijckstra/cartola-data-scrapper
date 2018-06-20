package request

import (
	"log"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dijckstra/cartola-data-scrapper/data/database"
	"github.com/dijckstra/cartola-data-scrapper/data/model"
)

const (
	playersEndpoint = "https://api.cartolafc.globo.com/atletas/mercado"
)

// PlayerRequestor is the object that
// retrieves player information from the API.
type PlayerRequestor struct {
	Db *database.DB
}

// RequestPlayers retrieves all players in the tournament.
func (requestor *PlayerRequestor) RequestPlayers() {
	// request rounds list
	resp, err := http.Get(playersEndpoint)
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
	var objmap map[string]*json.RawMessage
	err = json.Unmarshal(body, &objmap)

	var res []model.Player
	err = json.Unmarshal(*objmap["atletas"], &res)
	if err != nil {
		panic(err)
	}

	// insert into database
	err = requestor.Db.InsertPlayers(&res)
	if err != nil {
		panic(err)
	}
	log.Printf("sucesso")
}
