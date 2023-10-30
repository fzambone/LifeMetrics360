package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

const (
	DB_NAME                  = "life_metrics_360"
	EXPENSES_COLLECTION_NAME = "expenses"
)

// DatabaseMiddleware sets up a Gin middleware that attaches the DB connection to the context of every incoming request.
func DatabaseMiddleware(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dbClient", client)
		c.Next()
	}
}

// ConnectToDatabase connects to the MongoDB database
func ConnectToDatabase() (*mongo.Client, error) {
	//connectionString := "mongodb+srv://user_life_metrics_360:375au8uIvnreuWLw@lifemetrics360.dobldzn.mongodb.net/?retryWrites=true&w=majority"
	connectionString := os.Getenv("MONGO_CONNECTION_STRING")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10 seconds timeout
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil, err
	}

	return client, nil
}

// Close closes the MongoDB database connection
func Close(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10 seconds timeout
	defer cancel()
	err := client.Disconnect(ctx)
	if err != nil {
		log.Printf("Error closing database connection: %v", err)
		return err
	}

	return nil
}

// InsertExpense inserts a new expense record into the database.
func InsertExpense(db *mongo.Client, expense Expense) error {
	collection := db.Database(DB_NAME).Collection(EXPENSES_COLLECTION_NAME)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10 seconds timeout
	defer cancel()

	_, err := collection.InsertOne(ctx, expense)
	if err != nil {
		log.Printf("Error inserting expense record: %v", err)
		return err
	}

	return nil
}

// FindAllExpenses gets all the expense records from the database
func FindAllExpenses(client *mongo.Client) ([]Expense, error) {
	collection := client.Database(DB_NAME).Collection(EXPENSES_COLLECTION_NAME)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10 seconds timeout
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		log.Printf("Error getting expense records: %v", err)
		return nil, err
	}

	var expenses []Expense
	for cursor.Next(ctx) {
		var expense Expense
		err := cursor.Decode(&expense)
		if err != nil {
			log.Printf("Error decoding expense record: %v", err)
			return nil, err
		}

		expenses = append(expenses, expense)
	}

	return expenses, nil
}
