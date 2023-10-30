package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type Expense struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Amount   float64            `json:"amount"`
	Category string             `json:"category"`
	Merchant string             `json:"merchant"`
	Date     string             `json:"date"`
}

// Custom JSON Marshaling for Expense Struct
func (e *Expense) MarshalJSON() ([]byte, error) {
	type Alias Expense
	return json.Marshal(&struct {
		ID string `json:"id"`
		*Alias
	}{
		ID:    e.ID.Hex(),
		Alias: (*Alias)(e),
	})
}

// Custom JSON Unmarshaling for Expense Struct
func (e *Expense) UnmarshalJSON(data []byte) error {
	type Alias Expense
	aux := &struct {
		ID string `json:"id"`
		*Alias
	}{
		Alias: (*Alias)(e),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	e.ID, _ = primitive.ObjectIDFromHex(aux.ID)
	return nil
}

func RegisterRoutes(r *gin.Engine) {
	// Open DB connection and attaches it using the middleware
	//re := gin.Default()
	dbClient, _ := ConnectToDatabase()
	r.Use(DatabaseMiddleware(dbClient))

	// Define API endpoints here
	r.GET("/expenses", GetExpenses)
	r.POST("/expenses", CreateExpense)
}

func CreateExpense(c *gin.Context) {
	// Retrieve database connection
	db, _ := c.MustGet("dbClient").(*mongo.Client)

	// Decode the JSON request body
	var expense Expense
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&expense)
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
	// Retrieve database connection
	db, _ := c.MustGet("dbClient").(*mongo.Client)

	// Get the user's expenses
	expenses, err := FindAllExpenses(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Marshal the expenses to JSON
	expensesJSON, err := json.Marshal(expenses)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Decode the JSON expenses
	decodedExpenses, err := DecodeJSONExpenses(expensesJSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
		return
	}

	//Respond to the request
	c.JSON(http.StatusOK, decodedExpenses)
}

func DecodeJSONExpenses(expensesJSON []byte) ([]Expense, error) {
	var expenses []Expense
	err := json.Unmarshal(expensesJSON, &expenses)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}
