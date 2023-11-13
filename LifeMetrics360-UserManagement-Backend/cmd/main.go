package cmd

import (
	"github.com/fzambone/LifeMetrics360-UserManagement/handlers"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	//if err := godotenv.Load(); err != nil {
	//	logrus.Fatal("Error loading .env file")
	//}
	//
	//mongoURI := os.Getenv("MONGO_URI")

	// Initialize Echo
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("mysecret"),
	}))

	// Routes
	e.POST("/users", handlers.CreateUser)
	e.GET("/users:id", handlers.GetUser)
	e.PUT("/users:id", handlers.UpdateUser)
	e.DELETE("/users:id", handlers.DeleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":0881"))
}
