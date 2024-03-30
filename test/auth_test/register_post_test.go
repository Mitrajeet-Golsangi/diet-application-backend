package auth_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/app/auth"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/models"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/test/testdata"
	"github.com/stretchr/testify/assert"
)

func TestRegisterPost(t *testing.T) {
	r := testdata.SetupRouter()
	DB, mock := testdata.InitializeTestDB()

	sampleData := models.User{
		Name:        "AAA",
		Email:       "aaa@test.com",
		Username:    "aaa",
		Password:    "password",
		PhoneNumber: 1234567890,
		Gender:      "Male",
	}

	jsonBody, _ := json.Marshal(sampleData)

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO \"users\" (.+) VALUES (.+)").WillReturnRows(mock.NewRows([]string{"id"}).AddRow("1"))
	mock.ExpectCommit()

	r.POST("/", auth.RegisterPost(DB))
	
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonBody))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Nil(t, mock.ExpectationsWereMet())
}
