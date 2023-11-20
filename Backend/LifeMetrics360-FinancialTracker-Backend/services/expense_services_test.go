package services

import (
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

import (
	"context"
	"github.com/fzambone/LifeMetrics360-FinancialTracker/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ctx = context.TODO()

func TestExpenseService_InsertExpense(t *testing.T) {
	mockDB := new(MockDB)
	service := NewExpenseService(mockDB)

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
	mockDB := new(MockDB)
	service := NewExpenseService(mockDB)

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

	// Create the MockCursor with expected expenses
	mockCursor := new(MockCursor)
	mockCursor.Data = make([]interface{}, len(expectedExpenses))
	for i, v := range expectedExpenses {
		mockCursor.Data[i] = v
	}

	mockCursor.On("All", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*[]models.Expense)
		*arg = expectedExpenses
	}).Return(nil)

	// Set up the mock expectations
	mockDB.On("Find", context.Background(), expenseCollection, bson.D{}).Return(mockCursor, nil)

	// Test the FindAllExpenses function
	expenses, err := service.FindAllExpenses()

	// Assert that the expectations are met
	assert.NoError(t, err)
	assert.Equal(t, expectedExpenses, expenses)

	// Verify that ALL was called on the cursor
	mockCursor.AssertExpectations(t)
}

func TestExpenseService_UpdateExpense(t *testing.T) {
	mockDB := new(MockDB)
	service := NewExpenseService(mockDB)

	mockUpdatedResult := new(MockUpdateResult)
	mockUpdatedResult.On("MatchedCount").Return(int64(1))
	mockUpdatedResult.On("ModifiedCount").Return(int64(1))
	mockUpdatedResult.On("UpsertedCount").Return(int64(0))
	mockUpdatedResult.On("UpsertedID").Return(nil)
	expenseID := primitive.NewObjectID()
	updatedExpense := models.Expense{
		Amount:   200,
		Category: "Entertainment",
		Merchant: "Movie Theater",
		Date:     "2023-11-05",
	}

	// Set up the mock expectations
	mockDB.On("UpdateOne", ctx, expenseCollection, bson.M{"_id": expenseID}, bson.M{"$set": updatedExpense}).Return(mockUpdatedResult, nil)

	// Test the UpdateExpense function
	err := service.UpdateExpense(ctx, expenseID, updatedExpense)

	// Assert that the expectations are met
	assert.NoError(t, err)
}
