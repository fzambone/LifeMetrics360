package api

import (
	"encoding/json"
	"errors"
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

func (e *Expense) Validate() error {
	if e.Merchant == "" {
		return errors.New("merchant cannot be empty")
	}
	if e.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}
	if e.Date == "" {
		return errors.New("date cannot be empty")
	}
	if e.Category == "" {
		return errors.New("category cannot be empty")
	}
	return nil
}

func RegisterRoutes(r *gin.Engine) {
	// Define API endpoints here
	r.GET("/expenses", GetExpenses)
	r.POST("/expenses", CreateExpense)
	r.PUT("/expenses/:id", UpdateExpense)
}

func UpdateExpense(c *gin.Context) {
	// Connect to Database
	db := ConnectToDatabase()
	defer func(client *mongo.Client) {
		Close(client)
	}(db)

	// Get te expense ID from the URL parameters
	expenseID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		logrus.WithError(err).Error("Invalid ObjectID format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing expense ID"})
		return
	}

	// Decode the JSON request body
	var updatedExpense Expense
	if err := c.ShouldBindJSON(&updatedExpense); err != nil {
		Log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error decoding JSON request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the updated expense data
	if err := updatedExpense.Validate(); err != nil {
		Log.WithError(err).Error("Error validating expense")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the expense in the database
	err = UpdateExpenseInDB(db, expenseID, updatedExpense)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error updating expense in database")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update successful
	Log.WithFields(logrus.Fields{
		"expenseID": expenseID,
	}).Info("Expense updated successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Expense updated successfully"})
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
