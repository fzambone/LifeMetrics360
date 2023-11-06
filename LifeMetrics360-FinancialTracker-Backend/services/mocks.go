package services

import (
	"context"
	"github.com/fzambone/LifeMetrics360-FinancialTracker/utils"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MockDB is a mock type for the database interactions
type MockDB struct {
	mock.Mock
}

// InsertOne mocks the InsertOne method of the DatabaseHelper interface
func (m *MockDB) InsertOne(ctx context.Context, collection string, document interface{}) (primitive.ObjectID, error) {
	args := m.Called(ctx, collection, document)
	return args.Get(0).(primitive.ObjectID), args.Error(1)
}

// Find mocks the Find method of the DatabaseHelper interface
func (m *MockDB) Find(ctx context.Context, collection string, filter interface{}) (utils.CursorHelper, error) {
	args := m.Called(ctx, collection, filter)
	return args.Get(0).(utils.CursorHelper), args.Error(1)
}

// UpdateOne mocks the UpdateOne method of the DatabaseHelper interface
func (m *MockDB) UpdateOne(ctx context.Context, collection string, filter interface{}, update interface{}) (utils.UpdateResultHelper, error) {
	args := m.Called(ctx, collection, filter, update)
	return args.Get(0).(utils.UpdateResultHelper), args.Error(1)
}
