package token

import (
	"log"
	"strings"

	firebaseapp "github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/firebase_app"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/models"
	"github.com/gin-gonic/gin"
)

// Check if the token provided in the request is valid
func TokenValid(c *gin.Context) error {
	// Extract the token from the request
	tokenString := ExtractToken(c)

	// Check if the token is valid
	token, err := models.AuthClient.VerifyIDToken(firebaseapp.Ctx, tokenString)
	if err != nil {
		log.Println("Error verifying ID token: %v\n", err)
		return err
	}
	log.Println("Verified ID token: %v\n", token)
	return nil
}

// Extract the token from the request
func ExtractToken(c *gin.Context) string {
	// Extract the token from the query parameters or the Authorization header
	token := c.Query("token")

	// If the token is not empty return it
	if token != "" {
		return token
	}

	// If the token is empty, extract it from the Authorization header
	bearerToken := c.Request.Header.Get("Authorization")

	// If the token is not empty return it
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	// Return an empty string if the token is not found
	return ""
}

// Extract the user ID from the token
func GetUserID(c *gin.Context) (string, error) {
	// Extract the token from the request
	tokenString := ExtractToken(c)

	// Verify the token
	token, err := models.AuthClient.VerifyIDToken(firebaseapp.Ctx, tokenString)

	// Throw an error if the user is invalid
	if err != nil {
		log.Printf("Error verifying ID token while extracting User ID: %v\n", err)
		return "", err
	}

	// Return the user ID if the token is valid
	log.Println("User ID successfully retrieved from the token")
	return token.UID, nil
}
