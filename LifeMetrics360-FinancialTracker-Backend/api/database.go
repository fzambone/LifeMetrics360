package api

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

// ConnectToDatabase connects to the MongoDB database
func ConnectToDatabase() *mongo.Client {

	connectionString := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error connecting to database")
		return nil
	}

	return client
}

// Close closes the MongoDB database connection
func Close(client *mongo.Client) {
	err := client.Disconnect(context.Background())
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error closing database connection: %v", err)
	}

	Log.Info("Successfully closed database connection")
}

// InsertExpense inserts a new expense record into the database.
func InsertExpense(db *mongo.Client, expense Expense) (string, error) {
	collection := db.Database("expenses").Collection("expenses")

	result, err := collection.InsertOne(context.Background(), expense)
	if err != nil {
		log.Printf("Error inserting expense record: %v", err)
		Log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error inserting expense record: %v", err)
		return "", err
	}

	//Convert the inserted ID to string
	insertedID := result.InsertedID.(primitive.ObjectID).Hex()
	Log.Info("Successfully inserted expense %v", insertedID)
	return insertedID, nil
}

// FindAllExpenses gets all the expense records from the database
func FindAllExpenses(client *mongo.Client) []Expense {

	collection := client.Database("life_metrics_360").Collection("expenses")

	cursor, err := collection.Find(context.Background(), bson.M{})

	// Ensure the client is connected
	if err := client.Ping(context.Background(), nil); err != nil {
		Log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error pinging database: %v", err)
		return nil
	} else {
		Log.Info("Database connection is working!")
	}

	if err != nil {
		Log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error getting expense records: %v", err)
		return nil
	}

	var expenses []Expense
	for cursor.Next(context.Background()) {
		var expense Expense
		err := cursor.Decode(&expense)
		if err != nil {
			Log.WithFields(logrus.Fields{
				"error": err,
			}).Error("Error decoding expense record: %v", err)
			return nil
		}
		Log.Info("Successfully decoded expense %v", expense)
		expenses = append(expenses, expense)
	}

	Log.Info("Successfully got all expenses")
	return expenses
}
