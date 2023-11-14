package main

import (
	"context"
	"github.com/fzambone/LifeMetrics360-UserManagement/handlers"
	"github.com/fzambone/LifeMetrics360-UserManagement/services"
	"github.com/fzambone/LifeMetrics360-UserManagement/utils"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	e := echo.New()

	// Load .env file
	if err := godotenv.Load(); err != nil {
		logrus.WithError(err).Error("Error loading .env file")
	}

	// Connect to MongoDB
	db, err := utils.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to the database 2: %v", err)
	}
	defer db.Close()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("mysecret"),
	}))

	// Initialize handlers with db connection
	userService := services.NewUserService(db)
	handlerInstance := handlers.NewHandlers(userService)

	// Routes
	e.POST("/users", handlerInstance.CreateUser)
	e.GET("/users:id", handlerInstance.GetUser)
	e.PUT("/users:id", handlerInstance.UpdateUser)
	e.DELETE("/users:id", handlerInstance.DeleteUser)

	// Start server
	go func() {
		if err := e.Start(":8081"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// Close database connection
	db.Close()

	// Shutdown the server
	if err := e.Shutdown(context.Background()); err != nil {
		e.Logger.Fatal(err)
	}
}
