package utils

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

type Database struct {
	Client *mongo.Client
}

// NewDatabase creates a new Database instance
func NewDatabase() (*Database, error) {
	uri := os.Getenv("MONGO_URI")

	// Set a 10 second timeout for the initial connection to the MongoDB server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Ensure the context is canceled once done.

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Confirm the connection is established
	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return &Database{Client: client}, nil
}

// Collection returns a handle to a MongoDB collection with the provided name
func (db *Database) Collection(collectionName string) *mongo.Collection {
	return db.Client.Database("financial_tracker").Collection(collectionName)
}

// Close handles closing the connection to the database.
func (db *Database) Close() {
	if db != nil {
		// Set a 5 second timeout for the initial connection to the MongoDB server
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := db.Client.Disconnect(ctx); err != nil {
			logrus.WithError(err).Error("Failed to close database connection")
		}
	}
}
