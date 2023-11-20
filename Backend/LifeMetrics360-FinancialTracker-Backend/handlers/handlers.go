package handlers

import (
	"github.com/fzambone/LifeMetrics360-FinancialTracker/services"
	"github.com/fzambone/LifeMetrics360-Utils-Backend/utils"
)

type Handlers struct {
	DB             *utils.Database
	ExpenseService *services.ExpenseService
}

// NewHandlers creates a new Handlers struct with the necessary services
func NewHandlers(db *utils.Database, expenseService *services.ExpenseService) *Handlers {
	return &Handlers{
		DB:             db,
		ExpenseService: expenseService,
	}
}
