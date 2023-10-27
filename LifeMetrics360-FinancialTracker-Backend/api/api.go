package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type Expense struct {
	ID       int     `json:"id"`
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
	Merchant string  `json:"merchant"`
	Date     string  `json:"date"`
}

func RegisterRoutes(r *gin.Engine) {
	// Define API endpoints here
	r.GET("/expenses", GetExpenses)
	r.POST("/expenses", CreateExpense)
}

func CreateExpense(c *gin.Context) {
	// Connect to database
	db, err := ConnectToDatabase()
	if err != nil {
		log.Printf("Error connecting to database %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Close connection database before returning
	defer func(client *mongo.Client) {
		err := Close(client)
		if err != nil {
			log.Printf("Error closing database connection: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}(db)

	// Decode the JSON request body
	var expense Expense
	decoder := json.NewDecoder(c.Request.Body)
	err = decoder.Decode(&expense)
	if err != nil {
		log.Printf("Error decoding JSON request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the Expense
	// Check if all the required fields are present
	if expense.Amount == 0 || expense.Category == "" {
		log.Printf("Error validating expense: missing required fields")
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required fields"})
		return
	}

	// Check if the expense amount is valid
	if expense.Amount < 0 {
		log.Printf("Error validating expense: invalid amount")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid amount"})
		return
	}

	// Create the expense record
	err = InsertExpense(db, expense)
	if err != nil {
		log.Printf("Error inserting expense record: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Encode the expense as JSON
	expenseJSON, err := json.Marshal(expense)
	if err != nil {
		log.Printf("Error encoding expense as JSON: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond to the request
	c.JSON(http.StatusCreated, expenseJSON)
}

func GetExpenses(c *gin.Context) {
	// Connect to database
	db, err := ConnectToDatabase()
	if err != nil {
		log.Printf("Error connecting to database %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Close connection database before returning
	defer func(client *mongo.Client) {
		err := Close(client)
		if err != nil {
			log.Printf("Error closing database connection: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}(db)

	// Get the user's expenses
	expenses, err := FindAllExpenses(db)
	if err != nil {
		log.Printf("Error getting expenses: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Encode the expenses as JSON
	expensesJSON, err := json.Marshal(expenses)
	if err != nil {
		log.Printf("Error encoding expenses as JSON: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond to the request
	c.JSON(http.StatusOK, expensesJSON)
}
