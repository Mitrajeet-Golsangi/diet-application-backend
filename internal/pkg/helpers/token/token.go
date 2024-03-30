package token

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Generate a new JWT token after the user logs in
func GenerateToken(user_id uint) (string, error) {

	// Set the token lifespacn to 1 hour by default
	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "", err
	}

	// Create a new JWT token with the user ID and expiration time claims
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the API secret key and return the signed token
	return token.SignedString([]byte(os.Getenv("API_SECRET")))

}

// Check if the token provided in the request is valid
func TokenValid(c *gin.Context) error {
	// Extract the token from the request
	tokenString := ExtractToken(c)

	// Parse the token and check if it is valid
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
	
		return []byte(os.Getenv("API_SECRET")), nil
	})

	// throw the error if something unexpected went wrong while verifying the token
	if err != nil {
		return err
	}

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
func ExtractTokenID(c *gin.Context) (uint, error) {
	// Extract the token from the request
	tokenString := ExtractToken(c)

	// Parse the token and check if it is valid
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("API_SECRET")), nil
	})

	// throw the error if something unexpected went wrong while verifying the token
	if err != nil {
		return 0, err
	}

	// Get the claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)

	// Throw an error if the claims are not valid
	if ok && token.Valid {

		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		
		if err != nil {
			return 0, err
		}
		
		// Return the user ID if the claims are valid
		return uint(uid), nil
	}

	// Return 0 if the above logic fails to identify the token or throw and unexpected error
	return 0, nil
}
