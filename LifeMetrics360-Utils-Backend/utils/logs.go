package utils

import "github.com/gin-gonic/gin"

// HandleError takes care of sending a uniform error message
func HandleError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message})
	c.AbortWithStatus(statusCode)
}

// HandleInfo takes care of sending a uniform info message
//func HandleInfo(c *gin.Context, statusCode int, message string) {
//	c.JSON(statusCode, gin.H{"info": message})
//}
