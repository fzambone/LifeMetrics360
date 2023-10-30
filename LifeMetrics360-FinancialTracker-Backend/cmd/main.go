package main

import (
	"context"
	"github.com/fzambone/LifeMetrics360-FinancialTracker/api"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := gin.Default()

	// Register endpoint routes
	api.RegisterRoutes(r)

	// Gets a database connection
	dbClient, err := api.ConnectToDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create a channel to listen for interrupt signals
	quit := make(chan os.Signal, 1)

	// Notify the 'quit' channel for SIGINT and SIGTERM signals
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start the Gin server
	go func() {
		if err := r.Run(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Wait for an interrupt signal
	<-quit

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Disconnect from the database
	if err := dbClient.Disconnect(ctx); err != nil {
		log.Fatalf("Failed to gracefully disconnect from database: %v", err)
	}

	log.Println("Database connection closed")
	os.Exit(0)
}
