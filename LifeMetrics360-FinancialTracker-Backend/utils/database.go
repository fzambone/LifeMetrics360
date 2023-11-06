package utils

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

var _ DatabaseHelper = &Database{}

// CursorHelper abstracts the methods used from a database query result cursor
type CursorHelper interface {
	All(ctx context.Context, results interface{}) error
	Close(ctx context.Context) error
}

// UpdateResultHelper abstracts the methods used from the result of a database update operation.
type UpdateResultHelper interface {
	MatchedCount() int64
	ModifiedCount() int64
	UpsertedID() interface{}
}

type DatabaseHelper interface {
	InsertOne(ctx context.Context, collection string, document interface{}) (primitive.ObjectID, error)
	Find(ctx context.Context, collection string, filter interface{}) (CursorHelper, error)
	UpdateOne(ctx context.Context, collection string, filter interface{}, update interface{}) (UpdateResultHelper, error)
}

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
		// Set a 5 seconds timeout for the initial connection to the MongoDB server
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := db.Client.Disconnect(ctx); err != nil {
			logrus.WithError(err).Error("Failed to close database connection")
		}
	}
}

// InsertOne implements the DatabaseHelper interface method for inserting documents
func (db *Database) InsertOne(ctx context.Context, collection string, document interface{}) (primitive.ObjectID, error) {
	coll := db.Client.Database("financial_tracker").Collection(collection)

	result, err := coll.InsertOne(ctx, document)
	if err != nil {
		return primitive.NilObjectID, err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, errors.New("invalid ObjectID in inserted ID")
	}

	return oid, nil
}

// Find implements the DatabaseHelper interface method for finding documents
func (db *Database) Find(ctx context.Context, collection string, filter interface{}) (CursorHelper, error) {
	//TODO: implement logic
	return nil, nil
}

// UpdateOne implements the DatabaseHelper interface method for finding documents
func (db *Database) UpdateOne(ctx context.Context, collection string, filter interface{}, update interface{}) (UpdateResultHelper, error) {
	//TODO: implement logic
	return nil, nil
}
