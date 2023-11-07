package api

import (
	"github.com/fzambone/LifeMetrics360-FinancialTracker/handlers"
	"github.com/fzambone/LifeMetrics360-FinancialTracker/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, h *handlers.Handlers) {
	// Define middlewares
	router.Use(middleware.ErrorHandlingMiddleware)
	router.Use(middleware.InfoLoggingMiddleware())

	// Define API endpoints here
	router.GET("/expenses", h.GetAllExpenses)
	router.POST("/expenses", h.CreateExpense)
	router.PUT("/expenses/:id", h.UpdateExpense)
	router.DELETE("/expenses/:id", h.DeleteExpense)
}
