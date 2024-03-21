package auth

import (
	"github.com/gin-gonic/gin"
)

// PingGet returns a Gin handler function that handles the GET request for the /ping path.
// The handler function returns a JSON response with the message "pong".
func PingGet() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}
