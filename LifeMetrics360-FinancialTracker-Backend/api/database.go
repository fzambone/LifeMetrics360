package api

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// ConnectToDatabase connects to the MongoDB database
func ConnectToDatabase() (*mongo.Client, error) {
	// TODO: Fix connection string for mongoDB
	connectionString := "mongodb+srv://user_life_metrics_360:375au8uIvnreuWLw@cluster0.lsw8mua.mongodb.net/?retryWrites=true&w=majority"

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil, err
	}

	return client, nil
}

// Close closes the MongoDB database connection
func Close(client *mongo.Client) error {
	err := client.Disconnect(context.Background())
	if err != nil {
		log.Printf("Error closing database connection: %v", err)
		return err
	}

	return nil
}

// InsertExpense inserts a new expense record into the database.
func InsertExpense(db *mongo.Client, expense Expense) error {
	collection := db.Database("expenses").Collection("expenses")

	_, err := collection.InsertOne(context.Background(), expense)
	if err != nil {
		log.Printf("Error inserting expense record: %v", err)
		return err
	}

	return nil
}

// FindAllExpenses gets all the expense records from the database
func FindAllExpenses(db *mongo.Client) ([]Expense, error) {
	collection := db.Database("expenses").Collection("expenses")

	cursor, err := collection.Find(context.Background(), nil)
	if err != nil {
		log.Printf("Error getting expense records: %v", err)
		return nil, err
	}

	var expenses []Expense
	for cursor.Next(context.Background()) {
		var expense Expense
		err := cursor.Decode(&expenses)
		if err != nil {
			log.Printf("Error decoding expense record: %v", err)
			return nil, err
		}

		expenses = append(expenses, expense)
	}

	return expenses, nil
}
