package testdata

import "github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/helpers/token"

// Get the token for the user with provided ID in the test environment
func GetTestToken(uid uint) (string, error) {
	return token.GenerateToken(uid)
}