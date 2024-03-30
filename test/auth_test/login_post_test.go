package auth_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/app/auth"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/test/testdata"
	"github.com/stretchr/testify/assert"
)

func TestLoginPost(t *testing.T) {
	// Set the environment variables
	t.Setenv("TOKEN_HOUR_LIFESPAN", "1")
	t.Setenv("API_SECRET", "secret")

	// Setup the router and initialize the test database
	r := testdata.SetupRouter()
	DB, mock := testdata.InitializeTestDB()
	
	// Expect the select query to be executed and return the user data
	expectedQuery := "SELECT (.+) FROM \"users\" WHERE username = (.+) AND \"users\".\"deleted_at\" IS NULL LIMIT (.+)"
	expectedResult := mock.
		NewRows([]string{"id", "name", "email", "username", "password", "phone_number", "gender"}).
		AddRow(1, "John Doe", "john.doe@example.com", "johndoe", "password", 1234567890, "Male")

	mock.ExpectQuery(expectedQuery).WillReturnRows(expectedResult)

	// Mock the endpoint mapping to the login post request
	r.POST("/", auth.LoginPost(DB))
	
	// Send the login post request
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(`{"username": "johndoe", "password": "password"}`)))
	w := httptest.NewRecorder()
	
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, mock.ExpectationsWereMet())
}
