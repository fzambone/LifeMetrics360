package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type Expense struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Amount   float64            `json:"amount"`
	Category string             `json:"category"`
	Merchant string             `json:"merchant"`
	Date     string             `json:"date"`
}

func RegisterRoutes(r *gin.Engine) {
	// Define API endpoints here
	r.GET("/expenses", GetExpenses)
	r.POST("/expenses", CreateExpense)
}

func CreateExpense(c *gin.Context) {
	// Connect to database
	db := ConnectToDatabase()

	// Close connection database before returning
	defer func(client *mongo.Client) {
		Close(client)
	}(db)

	// Decode the JSON request body
	var expense Expense
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&expense)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error decoding JSON request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the Expense
	// Check if all the required fields are present
	if expense.Amount == 0 || expense.Category == "" {
		Log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error validating expense: missing required fields")
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required fields"})
		return
	}
	Log.Info("Expense validated, all required fields are present")

	// Check if the expense amount is valid
	if expense.Amount < 0 {
		Log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error validating expense: invalid amount")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid amount"})
		return
	}

	// Create the expense record
	insertedID, err := InsertExpense(db, expense)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error inserting expense record: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Encode the expense as JSON
	expenseJSON, err := json.Marshal(expense)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error encoding expense as JSON: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond to the request
	Log.Info("Expense %v created successfully", insertedID)
	c.JSON(http.StatusCreated, expenseJSON)
}

func GetExpenses(c *gin.Context) {
	// Connect to database
	db := ConnectToDatabase()

	// Close connection database before returning
	defer func(client *mongo.Client) {
		Close(client)
	}(db)

	// Get the user's expenses
	expenses := FindAllExpenses(db)
	if expenses == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error getting all expenses": nil})
		return
	}

	// Encode the expenses as JSON
	//expensesJSON, err := json.Marshal(expenses)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}

	//Respond to the request
	c.JSON(http.StatusOK, expenses)
}
