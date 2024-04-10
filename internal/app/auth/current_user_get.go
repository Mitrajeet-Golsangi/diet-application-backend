package auth

import (
	"net/http"

	firebaseapp "github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/firebase_app"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/helpers"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/helpers/token"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/models"
	"github.com/gin-gonic/gin"
)

// Get the information about the current logged in user using the JWT token
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, nil)

		// Extract the user ID from the token
		userID, err := token.GetUserID(c)

		if err != nil {
			helpers.BadRequest(c, err.Error())
			return
		}

		// Fetch the user from the database
		user, err := models.AuthClient.GetUser(firebaseapp.Ctx, userID)
		if err != nil {
			helpers.NotFound(c, err.Error())
			return
		}
		
		// Fetch the user health information from the database
		healthInfo, err := models.UserCollection.Doc(userID).Get(firebaseapp.Ctx)
		if err != nil {
			helpers.NotFound(c, err.Error())
			return
		}

		// Return the user
		c.JSON(http.StatusOK, gin.H{
			"message": "Success !",
			"user": user,
			"healthInfo": healthInfo,
		})
	}
}