package auth_test

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/app/auth"
// 	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/models"
// 	"github.com/Mitrajeet-Golsangi/diet-application-backend/test/testdata"
// 	"github.com/stretchr/testify/assert"
// )

// func TestRegisterPostSuccess(t *testing.T) {
// 	r := testdata.SetupRouter()

// 	// Setup the firestore database
// 	client := testdata.InitializeTestDB(context.Background())
// 	defer client.Close()

// 	// Create a sample user to insert in the database
// 	sampleData := models.User{
// 		Name:        "AAA",
// 		Email:       "aaa@test.com",
// 		Username:    "aaa",
// 		Password:    "password",
// 		PhoneNumber: 1234567890,
// 		Gender:      "Male",
// 	}

// 	// Convert the sample data to JSON
// 	jsonBody, _ := json.Marshal(sampleData)

// 	// Mock the endpoint mapping to the register post request
// 	r.POST("/", auth.RegisterPost())

// 	// Send the register post request
// 	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonBody))
// 	w := httptest.NewRecorder()

// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusCreated, w.Code)
// }
