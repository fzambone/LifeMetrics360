package cmd

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGO_URI")

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})

	e.Logger.Fatal(e.Start(":0881"))
}
