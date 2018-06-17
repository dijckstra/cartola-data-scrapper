package database

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
)

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
