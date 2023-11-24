package utils

import "github.com/gin-gonic/gin"

// RespondJSON sends a JSON response with a specific status code
func RespondJSON(c *gin.Context, status int, payload interface{}) {
	c.JSON(status, payload)
}

// RespondError sends a JSON error response with a specific status code and message
func RespondError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}
