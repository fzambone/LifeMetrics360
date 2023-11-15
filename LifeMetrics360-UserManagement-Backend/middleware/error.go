package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// ErrorLoggingMiddleware catches an error and logs it using logrus
func ErrorLoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			c.Error(err)
			logrus.WithError(err).Error("Error encountered")
		}
		return err
	}
}
