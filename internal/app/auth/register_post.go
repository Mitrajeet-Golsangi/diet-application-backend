package auth

import (
	"net/http"

	"firebase.google.com/go/v4/auth"
	firebaseapp "github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/firebase_app"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/helpers"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserRegistrationInput struct {
	Email         string `json:"email" binding:"required,email"`    // Email address of the user
	EmailVerified bool   `json:"email_verified"`                    // Whether the email address has been verified or not
	Password      string `json:"password" binding:"required,min=6"` // Password of the user
	DisplayName   string `json:"display_name" binding:"required"`   // Display name of the user
	PhoneNumber   string `json:"phone_number"`                      // Phone number of the user
	Disabled      bool   `json:"disabled"`                          // Whether the user account is disabled or not
}

// PingGet returns a Gin handler function that handles the GET request for the /ping path.
// The handler function returns a JSON response with the message "pong".
// ? Note: The Firebase cloud function will automatically create a new
// ? user information document in the Firestore database when a new user is registered.
func RegisterPost() gin.HandlerFunc {

	return func(c *gin.Context) {

		// Parse the request body into the User struct
		var u *UserRegistrationInput

		// Throw an error if the request body is invalid
		if err := c.ShouldBindJSON(&u); err != nil {
			helpers.BadRequest(c, err.Error())
			return
		}

		// Create a user using all the parameters passed in the request body
		params := (&auth.UserToCreate{}).
			Email(u.Email).
			EmailVerified(u.EmailVerified).
			Password(u.Password).
			DisplayName(u.DisplayName).
			PhoneNumber(u.PhoneNumber).
			Disabled(u.Disabled)

		// Create a new user in the database
		user, err := models.AuthClient.CreateUser(firebaseapp.Ctx, params)

		// Throw an error if the user creation fails
		if err != nil {
			helpers.InternalServerError(c, "Failed to create user ! "+err.Error())
			return
		}
		logrus.Info("New User Registered Successfully !")

		c.JSON(http.StatusCreated, user)
	}
}
