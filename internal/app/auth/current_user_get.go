package auth

import (
	"net/http"

	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/helpers"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/helpers/token"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/models"
	"github.com/gin-gonic/gin"
)

// Get the information about the current logged in user using the JWT token
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the user ID from the token
		userID, err := token.ExtractTokenID(c)

		if err != nil {
			helpers.BadRequest(c, err.Error())
			return
		}

		// Fetch the user from the database
		user, err := models.GetUserByID(userID)

		if err != nil {
			helpers.NotFound(c, err.Error())
			return	
		}

		// Return the user
		c.JSON(http.StatusOK, gin.H{
			"message": "Success !",
			"user": user,
		})
	}
}