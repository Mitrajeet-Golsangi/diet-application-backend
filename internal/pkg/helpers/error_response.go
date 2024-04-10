// This package is responsible for providing an error response to the client in case of an error.
package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Return a 400 Bad Request error to the client
func BadRequest(c *gin.Context, message string) {
	logrus.Error(message)
	c.JSON(http.StatusBadRequest, gin.H{
		"error": message,
	})
}

// Return a 401 Unauthorized error to the client
func Unauthorized(c *gin.Context, message string) {
	logrus.Error(message)
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": message,
	})
}

// Return a 403 Forbidden error to the client
func Forbidden(c *gin.Context, message string) {
	logrus.Error(message)
	c.JSON(http.StatusForbidden, gin.H{
		"error": message,
	})
}

// Return a 404 Not Found error to the client
func NotFound(c *gin.Context, message string) {
	logrus.Error(message)
	c.JSON(http.StatusNotFound, gin.H{
		"error": message,
	})
}

// Return a 405 Method Not Allowed error to the client
func MethodNotAllowed(message string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": message,
		})
	}
}

// Return a 500 Internal Server Error to the client
func InternalServerError(c *gin.Context, message string) {
	logrus.Error(message)
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": message,
	})
}

// Return a 501 Not Implemented error to the client
func NotImplemented(c *gin.Context, message string) {
	logrus.Error(message)
	c.JSON(http.StatusNotImplemented, gin.H{
		"error": message,
	})
}
