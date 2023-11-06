package services

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

import (
	"context"
	"github.com/fzambone/LifeMetrics360-FinancialTracker/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mockDB = new(MockDB)
var service = NewExpenseService(mockDB)
var ctx = context.TODO()

func TestExpenseService_InsertExpense(t *testing.T) {

	// create a new expense and the expected ID that would be returned
	expense := models.Expense{
		Amount:   150,
		Category: "Groceries",
		Merchant: "Supermarket",
		Date:     "2023-10-31",
	}

	expectedID := primitive.NewObjectID()

	// Set  up the mock expectations
	mockDB.On("InsertOne", ctx, expenseCollection, mock.AnythingOfType("models.Expense")).Return(expectedID, nil)

	// Test the CreateExpense function
	id, err := service.InsertExpense(ctx, expense)

	// Assert that the expectations are met
	assert.NoError(t, err)
	assert.Equal(t, expectedID, id)
}

func TestExpenseService_FindAllExpenses(t *testing.T) {
	expectedExpenses := []models.Expense{
		{
			ID:       primitive.NewObjectID(),
			Amount:   100,
			Category: "Utilities",
			Merchant: "Electric Company",
			Date:     "2023-11-01",
		},
		{
			ID:       primitive.NewObjectID(),
			Amount:   50,
			Category: "Groceries",
			Merchant: "Local Market",
			Date:     "2023-11-02",
		},
	}

	// Set up the mock expectations
	mockDB.On("Find", mock.AnythingOfType("*context.emptyCtx"), expenseCollection, bson.D{}).Return(expectedExpenses, nil)

	// Test the FindAllExpenses function
	expenses, err := service.FindAllExpenses()

	// Assert that the expectations are met
	assert.NoError(t, err)
	assert.Equal(t, expectedExpenses, expenses)
}

func TestExpenseService_UpdateExpense(t *testing.T) {
	expenseID := primitive.NewObjectID()
	updatedExpense := models.Expense{
		Amount:   200,
		Category: "Entertainment",
		Merchant: "Movie Theater",
		Date:     "2023-11-05",
	}

	updateResult := &mongo.UpdateResult{
		MatchedCount:  1,
		ModifiedCount: 1,
		UpsertedCount: 0,
		UpsertedID:    nil,
	}

	// Set up the mock expectations
	mockDB.On("UpdateOne", ctx, expenseCollection, bson.M{"_id": expenseID}, bson.M{"$set": updatedExpense}).Return(updateResult, nil)

	// Test the UpdateExpense function
	err := service.UpdateExpense(ctx, expenseID, updatedExpense)

	// Assert that the expectations are met
	assert.NoError(t, err)
}
