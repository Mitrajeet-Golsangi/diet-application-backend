// Description: The file contains all the information regarding
// the responses given by the API for the respective endpoints.
package api

// Parameters taken in by the API for the User Registration endpoint
type UserRegistrationParams struct {
	Username string
	Password string
	FullName string
	Email string
}

// Response given by the API for the User Registration endpoint
type UserRegistrationResponse struct {
	// The HTTP response code
	// 200 for success
	// 400 for bad request
	// 500 for internal server error
	Code int

	// The new user created in the database
	User User
}

type CoinBalanceProgram struct {
	Username string
}

type CoinBalanceResponse struct {
	Code int
	Balance int64
}

type Error struct {
	Code int
	
}