package handlers

import (
	"github.com/fzambone/LifeMetrics360-Utils-Backend/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

// DeleteExpense handles the PUT method and updates an existing expense by ID
func (h *Handlers) DeleteExpense(c *gin.Context) {
	// Get the expense ID from the URL parameters
	expenseID, err := primitive.ObjectIDFromHex(c.Param("id"))
	log.Print(expenseID)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "invalid expense ID format")
		return
	}

	// Delete expense from database
	err = h.ExpenseService.DeleteExpense(c, expenseID)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "error deleting expense from database")
		return
	}

	c.Set("info", "Expense deleted successfully: "+expenseID.Hex())
	c.JSON(http.StatusOK, gin.H{"message": "Expense deleted successfully"})
}
