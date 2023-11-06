package handlers

import (
	"github.com/fzambone/LifeMetrics360-FinancialTracker/models"
	"github.com/fzambone/LifeMetrics360-FinancialTracker/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// UpdateExpense handles the PUT method and updates an existing expense by ID
func (h *Handlers) UpdateExpense(c *gin.Context) {
	// Get the expense ID from the URL parameters
	expenseID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "invalid expense ID format")
		return
	}

	// Decode JSON form request body
	var updatedExpense models.Expense
	if err := c.ShouldBindJSON(&updatedExpense); err != nil {
		utils.HandleError(c, http.StatusBadRequest, "error decoding JSON request body")
		return
	}

	// Validates expense fields
	if err := updatedExpense.Validate(); err != nil {
		utils.HandleError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Update the expense in the database
	err = h.ExpenseService.UpdateExpense(c.Request.Context(), expenseID, updatedExpense)
	if err != nil {
		if err.Error() == "no document found with the specified ID" {
			utils.HandleError(c, http.StatusNotFound, "no expense found with the given ID")
		} else {
			utils.HandleError(c, http.StatusInternalServerError, "error updating expense in database")
		}
		return
	}

	// Returns success response
	c.Set("info", "Expense updated successfully: "+expenseID.Hex())
	c.JSON(http.StatusOK, gin.H{"message": "Expense updated successfully"})
}
