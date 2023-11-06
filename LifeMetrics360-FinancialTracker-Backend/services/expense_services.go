package services

import (
	"context"
	"errors"
	"github.com/fzambone/LifeMetrics360-FinancialTracker/models"
	"github.com/fzambone/LifeMetrics360-FinancialTracker/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const expenseCollection = "expenses"

type ExpenseService struct {
	DB *utils.Database
}

func NewExpenseService(db *utils.Database) *ExpenseService {
	return &ExpenseService{DB: db}
}

// InsertExpense inserts a new expense into the database
func (s *ExpenseService) InsertExpense(ctx context.Context, expense models.Expense) (primitive.ObjectID, error) {
	res, err := s.DB.Collection(expenseCollection).InsertOne(ctx, expense)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

// FindAllExpenses retrieves all expenses from the database
func (s *ExpenseService) FindAllExpenses() ([]models.Expense, error) {
	var expenses []models.Expense
	cursor, err := s.DB.Collection(expenseCollection).Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.Background(), &expenses); err != nil {
		return nil, err
	}
	return expenses, nil
}

// UpdateExpense updates an existing expense by ID
func (s *ExpenseService) UpdateExpense(ctx context.Context, id primitive.ObjectID, updatedExpense models.Expense) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedExpense}
	result, err := s.DB.Collection(expenseCollection).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	// Verify if Expense exists
	if result.MatchedCount == 0 {
		return errors.New("no document found with the specified ID")
	}

	return nil
}