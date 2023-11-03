package main

import (
	"github.com/fzambone/LifeMetrics360-FinancialTracker/api"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// Initialize logrus
	api.InitLogger()

	// Read .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	api.RegisterRoutes(r)

	r.Run()
}
