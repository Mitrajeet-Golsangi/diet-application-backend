package auth_test

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/app/auth"
// 	"github.com/Mitrajeet-Golsangi/diet-application-backend/test/testdata"
// 	"github.com/stretchr/testify/assert"
// )

// func TestCurrentUserSuccess(t *testing.T) {
// 	// Set the environment variables
// 	testdata.InitializeEnv(t)

// 	// Setup the router and initialize the test database
// 	r := testdata.SetupRouter()

// 	// Mock the endpoint mapping to the current user get request
// 	r.GET("/", auth.CurrentUser())

// 	req, _ := http.NewRequest("GET", "/", nil)

// 	// Set header of the request to the authorized token
// 	token, _ := testdata.GetTestToken(1)

// 	req.Header.Set("Authorization", "Bearer " + token)

// 	w := httptest.NewRecorder()

// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)
// }

// func TestCurrentUserNotFound(t *testing.T) {
// 	// Set the environment variables
// 	testdata.InitializeEnv(t)

// 	// Setup the router and initialize the test database
// 	r := testdata.SetupRouter()
// 	mock := testdata.InitializeTestDB()

// 	// Expect the select query to be executed and return the user data
// 	expectedQuery := "SELECT * FROM \"users\" WHERE \"users\".\"id\" = .+ AND \"users\".\"deleted_at\" IS NULL ORDER BY \"users\".\"id\" LIMIT .+"
// 	expectedResult := mock.
// 		NewRows([]string{"id", "name", "email", "username", "password", "phone_number", "gender"}).
// 		AddRow(1, "John Doe", "john.doe@example.com", "johndoe", "password", 1234567890, "Male")

// 	mock.ExpectQuery(expectedQuery).WillReturnRows(expectedResult)

// 	// Mock the endpoint mapping to the current user get request
// 	r.GET("/", auth.CurrentUser())

// 	req, _ := http.NewRequest("GET", "/", nil)

// 	// Set header of the request to the authorized token
// 	token, _ := testdata.GetTestToken(2)

// 	req.Header.Set("Authorization", "Bearer " + token)

// 	w := httptest.NewRecorder()

// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusNotFound, w.Code)
// 	assert.Equal(t, `{"error":"User not Found"}`, w.Body.String())
// }