package main

import (
	"context"
	"github.com/fzambone/LifeMetrics360-UserManagement/handlers"
	customMiddleware "github.com/fzambone/LifeMetrics360-UserManagement/middleware"
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

	jwtSecret := os.Getenv("JWT_SECRET")

	// Connect to MongoDB
	db, err := utils.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to the database 2: %v", err)
	}
	defer db.Close()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Custom Middlewares
	e.Use(customMiddleware.ErrorLoggingMiddleware)
	e.Use(customMiddleware.InfoLoggingMiddleware)

	// Initialize handlers with db connection
	userService := services.NewUserService(db)
	avatarService := services.NewAvatarService(db)
	userHandlerInstance := handlers.NewUserHandlers(userService)
	avatarHandlerInstance := handlers.NewAvatarHandlers(avatarService)

	// Routes
	e.GET("/avatar/:email", avatarHandlerInstance.Avatar)
	e.POST("/login", userHandlerInstance.Login)
	e.POST("/users", userHandlerInstance.CreateUser, echojwt.WithConfig(echojwt.Config{SigningKey: []byte(jwtSecret)}))
	e.GET("/users/:id", userHandlerInstance.GetUser, echojwt.WithConfig(echojwt.Config{SigningKey: []byte(jwtSecret)}))
	e.PUT("/users/:id", userHandlerInstance.UpdateUser, echojwt.WithConfig(echojwt.Config{SigningKey: []byte(jwtSecret)}))
	e.DELETE("/users/:id", userHandlerInstance.DeleteUser, echojwt.WithConfig(echojwt.Config{SigningKey: []byte(jwtSecret)}))

	log.Println(jwtSecret)

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
