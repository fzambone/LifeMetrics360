package middleware

import "github.com/labstack/echo/v4"

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: Implement JWT calidation
		return next(c)
	}
}

func ErrorLoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: Implement error logging
		return next(c)
	}
}

func InfoLoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: Implement info logging
		return next(c)
	}
}
