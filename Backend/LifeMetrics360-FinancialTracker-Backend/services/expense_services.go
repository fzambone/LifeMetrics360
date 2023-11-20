package services

import (
	"context"
	"errors"
	"github.com/fzambone/LifeMetrics360-FinancialTracker/models"
	"github.com/fzambone/LifeMetrics360-Utils-Backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

const expenseCollection = "expenses"

type ExpenseService struct {
	DB utils.DatabaseHelper
}

func NewExpenseService(db utils.DatabaseHelper) *ExpenseService {
	return &ExpenseService{DB: db}
}

// InsertExpense inserts a new expense into the database
func (s *ExpenseService) InsertExpense(ctx context.Context, expense models.Expense) (primitive.ObjectID, error) {
	res, err := s.DB.InsertOne(ctx, expenseCollection, expense)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res, nil
}

// FindAllExpenses retrieves all expenses from the database
func (s *ExpenseService) FindAllExpenses() ([]models.Expense, error) {
	var expenses []models.Expense
	cursor, err := s.DB.Find(context.Background(), expenseCollection, bson.D{})
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
	result, err := s.DB.UpdateOne(ctx, expenseCollection, filter, update)
	if err != nil {
		return err
	}

	// Verify if Expense exists
	if result.MatchedCount() == 0 {
		return errors.New("no document found with the specified ID")
	}

	return nil
}

// DeleteExpense deletes an existing expense by ID
func (s *ExpenseService) DeleteExpense(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	log.Print(filter)
	result, err := s.DB.DeleteOne(ctx, expenseCollection, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount() == 0 {
		return errors.New("no document found with the specified ID")
	}

	return nil
}
