package models

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Expense struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Amount   float64            `json:"amount"`
	Category string             `json:"category"`
	Merchant string             `json:"merchant"`
	Date     string             `json:"date"`
}

func (e *Expense) Validate() error {
	// Validate required fields
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

	// Validate Date format, assuming "2023-10-31" is the right format
	_, err := time.Parse("2006-01-02", e.Date)
	if err != nil {
		return errors.New("data must be in the format YYYY-MM-DD")
	}

	return nil
}
