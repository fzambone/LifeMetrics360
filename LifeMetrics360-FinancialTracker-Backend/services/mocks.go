package services

import (
	"context"
	"github.com/fzambone/LifeMetrics360-Utils-Backend/utils"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
)

// MockDB is a mock type for the database interactions
type MockDB struct {
	mock.Mock
}

// MockCursor is a mock implementation of utils.CursorHelper for test purposes
type MockCursor struct {
	mock.Mock
	Data []interface{}
}

type MockUpdateResult struct {
	mock.Mock
}

type MockDeleteResult struct {
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

// All mock the All method of the CursorHelper interface
func (m *MockCursor) All(ctx context.Context, results interface{}) error {
	args := m.Called(ctx, results)
	resultVal := reflect.ValueOf(results).Elem()
	// Clear the existing slice before setting it with new data to avoid duplication
	if resultVal.Kind() == reflect.Slice {
		resultVal.Set(reflect.MakeSlice(resultVal.Type(), 0, len(m.Data)))
	}
	for _, d := range m.Data {
		resultVal.Set(reflect.Append(resultVal, reflect.ValueOf(d)))
	}
	return args.Error(0)
}

// Close mocks the Close method of the CursorHelper interface
func (m *MockCursor) Close(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockUpdateResult) MatchedCount() int64 {
	args := m.Called()
	return args.Get(0).(int64)
}

func (m *MockUpdateResult) ModifiedCount() int64 {
	args := m.Called()
	return args.Get(0).(int64)
}

func (m *MockUpdateResult) UpsertedCount() int64 {
	args := m.Called()
	return args.Get(0).(int64)
}

func (m *MockUpdateResult) UpsertedID() interface{} {
	args := m.Called()
	return args.Get(0)
}

func (m *MockDeleteResult) DeletedCount() int64 {
	args := m.Called()
	return args.Get(0).(int64)
}
