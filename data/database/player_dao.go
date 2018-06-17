package database

import (
	"context"

	"github.com/dijckstra/cartola-data-scrapper/data/model"
)

const (
	playersCollection = "players"
)

// InsertPlayers executes the bulk insertion of players.
func (db *DB) InsertPlayers(players *[]model.Player) error {
	// Required conversion to []interface{}
	s := make([]interface{}, len(*players))
	for i, v := range *players {
		s[i] = v
	}

	// TODO update collection instead of reinsertions
	_, err := db.Collection(playersCollection).DeleteMany(context.Background(), nil)
	_, err = db.Collection(playersCollection).InsertMany(context.Background(), s)
	return err
}
