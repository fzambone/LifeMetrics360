package handlers

import (
	"github.com/fzambone/LifeMetrics360-FinancialTracker/models"
	"github.com/fzambone/LifeMetrics360-FinancialTracker/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateExpense handles the POST request to create a new request
func (h *Handlers) CreateExpense(c *gin.Context) {
	// Decode the JSON request body
	var expense models.Expense
	if err := c.ShouldBindJSON(&expense); err != nil {
		utils.HandleError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate the expense data
	if err := expense.Validate(); err != nil {
		utils.HandleError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Create the expense record in the database
	insertedID, err := h.ExpenseService.InsertExpense(c.Request.Context(), expense)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Failed to create expense")
		return
	}

	// Respond to the request
	c.Set("info", "New expense created: "+insertedID.Hex())
	c.JSON(http.StatusCreated, gin.H{"id": insertedID})
}
