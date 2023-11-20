package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// InfoLoggingMiddleware
func InfoLoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		logrus.WithFields(logrus.Fields{
			"method": c.Request().Method,
			"path":   c.Request().URL.Path,
			"status": c.Response().Status,
		}).Info("Request processed")
		return err
	}
}
