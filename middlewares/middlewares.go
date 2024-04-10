package middlewares

import (
	"net/http"

	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/helpers/token"
	"github.com/gin-gonic/gin"
)

// Authenticate the user using the JWT token
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)

		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Next()
	}
}
