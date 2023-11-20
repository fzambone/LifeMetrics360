package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ErrorHandlingMiddleware catches any errors thrown in the handlers and logs them
// TODO: Improve error handling, too generic
func ErrorHandlingMiddleware(c *gin.Context) {
	c.Next() // execute all the handlers

	// Handle errors if there are any
	if len(c.Errors) > 0 {
		for _, e := range c.Errors {
			logrus.WithError(e.Err).Error("An error occurred")
		}
	}
}
