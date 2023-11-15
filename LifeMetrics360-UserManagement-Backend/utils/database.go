package utils

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

type Handlers struct {
	DB *Database
}

type Database struct {
	Client *mongo.Client
}

// Collection returns a handler to a specific collection in the MongoDB
func (db *Database) Collection(name string) *mongo.Collection {
	return db.Client.Database("users").Collection(name)
}

// NewDatabase creates a new database connection to be used
func NewDatabase() (*Database, error) {
	uri := os.Getenv("MONGO_URI")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		logrus.WithError(err).Error("Failed to connect to database 1")
		return nil, err
	}

	// Ping the database to verify connection is established
	if err := client.Ping(ctx, nil); err != nil {
		logrus.WithError(err).Error("Failed to ping database")
		return nil, err
	}

	logrus.Info("Connected to MongoDB")
	return &Database{Client: client}, nil
}

// Close closes the connection with the database
func (db *Database) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.Client.Disconnect(ctx); err != nil {
		logrus.WithError(err).Error("Failed to close database connection")
	} else {
		logrus.Info("Database connection closed")
	}
}
