package database

import (
	"context"

	"github.com/dijckstra/cartola-data-scrapper/data/model"
)

const (
	collection = "rounds"
)

// InsertRounds executes the bulk insertion of rounds.
func (db *DB) InsertRounds(rounds []model.Round) error {
	// Required conversion to []interface{}
	s := make([]interface{}, len(rounds))
	for i, v := range rounds {
		s[i] = v
	}

	_, err := db.Collection(collection).InsertMany(context.Background(), s)
	return err
}
