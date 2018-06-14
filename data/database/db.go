package database

import (
	"context"

	"github.com/dijckstra/cartola-data-scrapper/data/model"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Database is the interface to the operations done to MongoDB.
type Database interface {
	InsertRounds(rounds []*model.Round) error
}

// DB is our MongoDB wrapper.
type DB struct {
	*mongo.Database
}

// NewDB returns a MongoDB database.
func NewDB(server string, database string) (*DB, error) {
	client, err := mongo.Connect(context.Background(), server, nil)
	if err != nil {
		return nil, err
	}

	db := client.Database(database)
	return &DB{db}, nil
}
