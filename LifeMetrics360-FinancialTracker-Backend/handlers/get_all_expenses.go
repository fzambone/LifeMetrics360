package handlers

import (
	"github.com/fzambone/LifeMetrics360-FinancialTracker/models"
	"github.com/fzambone/LifeMetrics360-Utils-Backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllExpenses handles the GET request and retrieves all the expenses from the database and returns them
func (h *Handlers) GetAllExpenses(c *gin.Context) {
	// Get the user's expenses
	expenses, err := h.ExpenseService.FindAllExpenses()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Error retrieving all expenses from the database")
		return
	}

	// If no expense found, return an empty slice
	if expenses == nil {
		expenses = []models.Expense{}
	}

	// Response to the request with the retrieved expenses
	c.Set("info", "Expenses retrieved successfully")
	c.JSON(http.StatusOK, gin.H{"expenses": expenses})
}
