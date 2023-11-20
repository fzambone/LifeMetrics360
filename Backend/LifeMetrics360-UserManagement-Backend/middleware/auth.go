package middleware

//
//import (
//	"fmt"
//	"github.com/dgrijalva/jwt-go"
//	"github.com/labstack/echo/v4"
//	"log"
//	"net/http"
//	"os"
//	"strings"
//)
//
//// AuthMiddleware validates the JWT Token on every request in a route that is registered to ask for validation
//func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		// Extract the token from Authorization header
//		authHeader := c.Request().Header.Get("Authorization")
//		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
//		jwtSecret := os.Getenv("JWT_SECRET")
//		log.Print(jwtSecret)
//
//		if tokenString == "" {
//			return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
//		}
//
//		// Parse the token
//		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//			// Make sure token's method matches signing method
//			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//			}
//			return []byte(jwtSecret), nil
//		})
//
//		if err != nil {
//			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
//		}
//
//		if !token.Valid {
//			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
//		}
//
//		return next(c)
//	}
//}
