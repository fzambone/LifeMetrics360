package main

import (
	"github.com/fzambone/LifeMetrics360-FinancialTracker/api"
	"github.com/fzambone/LifeMetrics360-FinancialTracker/handlers"
	"github.com/fzambone/LifeMetrics360-FinancialTracker/services"
	"github.com/fzambone/LifeMetrics360-Utils-Backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// Initialize logrus
	api.InitLogger()

	// Read .env file
	if err := godotenv.Load(); err != nil {
		logrus.WithError(err).Fatal("Could not load env file")
	}

	// Open new database connection
	db, err := utils.NewDatabase()
	if err != nil {
		logrus.WithError(err).Fatal("Could not connect to the database")
	}
	defer db.Close()

	// Create an instance of ExpenseService
	expenseService := services.NewExpenseService(db)

	router := gin.Default()

	// Create an instance of Handlers with a db dependency
	h := handlers.NewHandlers(db, expenseService)

	// Register API routes
	api.RegisterRoutes(router, h)

	router.Run(":8080")
}
