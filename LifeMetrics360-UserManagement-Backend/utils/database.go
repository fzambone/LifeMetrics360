package utils

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

//const userCollection = "users"

type Handlers struct {
	DB *Database
}

type Database struct {
	Client *mongo.Client
}

// Collection returns a handle to a specific collection in the MongoDB
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

//// addUserToDb adds a new user to the database
//func (h *Handlers) addUserToDB(client *mongo.Client, user *models.User) (primitive.ObjectID, error) {
//	collection := client.Database("users").Collection(userCollection)
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	// Hash the password before storing in the database
//	user.Password, _ = utils.HashPassword(user.Password)
//	result, err := collection.InsertOne(ctx, user)
//	if err != nil {
//		return primitive.NilObjectID, err
//	}
//
//	return result.InsertedID.(primitive.ObjectID), nil
//}
//
//// getUserFromDB retrieves a user by their ID from database
//func (h *Handlers) getUserFromDB(client *mongo.Client, userID string) (*models.User, error) {
//	collection := client.Database("users").Collection(userCollection)
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	objID, _ := primitive.ObjectIDFromHex(userID)
//	var user models.User
//	err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
//	if err != nil {
//		return nil, err
//	}
//	return &user, nil
//}
//
//// updateUserInDB updates a user's details in the database
//func (h *Handlers) updateUserInDB(client *mongo.Client, userID string, updateData *models.User) error {
//	collection := client.Database("users").Collection(userCollection)
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	objID, _ := primitive.ObjectIDFromHex(userID)
//	_, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
//	return err
//}
