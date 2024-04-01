package auth

import (
	"log"
	"net/http"

	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/helpers"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/models"
	"github.com/gin-gonic/gin"
)

// PingGet returns a Gin handler function that handles the GET request for the /ping path.
// The handler function returns a JSON response with the message "pong".
func RegisterPost() gin.HandlerFunc {

	return func(c *gin.Context) {

		// Parse the request body into the User struct
		var user *models.User

		// Throw an error if the request body is invalid
		if err := c.ShouldBindJSON(&user); err != nil {
			helpers.BadRequest(c, err.Error())
			return
		}

		// Create a new user in the database
		user, err := user.Save()

		// Throw an error if the user creation fails
		if err != nil {
			helpers.InternalServerError(c, "Failed to create user !")
			return
		}

		// Throw an error if the JSON conversion fails
		if err != nil {
			helpers.InternalServerError(c, "Failed to convert user to JSON !")
			return
		}

		log.Println("New User Registered Successfully !")

		c.JSON(http.StatusCreated, user)
	}
}
