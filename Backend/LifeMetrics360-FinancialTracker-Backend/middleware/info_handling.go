package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// InfoLoggingMiddleware catches any successful messages thrown in the handlers and logs them
// Option One
//func InfoLoggingMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		c.Next()
//
//		// Logs status code if response code is between 200 and 300 range
//		if c.Writer.Status() >= 200 && c.Writer.Status() < 300 {
//			logrus.WithFields(logrus.Fields{
//				"status": c.Writer.Status(),
//				"method": c.Request.Method,
//				"path":   c.Request.URL.Path,
//			}).Info("Request processed successfully")
//		}
//	}
//}

// InfoLoggingMiddleware catches any info messages thrown in the handlers and logs them
// Option Two (preferred by me)
func InfoLoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Catches if any 'info' message exists
		if info, exists := c.Get("info"); exists {
			logrus.Info(info)
		}
	}
}
