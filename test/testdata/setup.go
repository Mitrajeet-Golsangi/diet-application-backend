package testdata

import (
	"context"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
	firebaseapp "github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/firebase_app"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/models"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// Run the Firestore emulator before running the tests
// Also teardown the emulator after the tests are finished

// Initialize the testing database with the required tables and sample data
// Users:
//
//	John Doe
//	Jane Doe
//
// Health Information:
//
//	John Doe => Weight: 70, Height: 180, BMI: 21.6, Birthday: 1990-01-01, Exercise Frequency: 3
//	Jane Doe => Weight: 80, Height: 160, BMI: 31.2, Birthday: 1990-02-01, Exercise Frequency: 3
//
// Exercise Information:
//
//	Running
//	Circuit training, minimal rest
//	Weight lifting, light workout
func InitializeTestDB(ctx context.Context) *firestore.Client {
	// Create a new Firestore test client
	firebaseapp.Initialize()
	models.InitDatabase()

	// Create sample new user data
	user1 := models.User{
		Name:        "John Doe",
		Email:       "john.doe@example.com",
		Username:    "johndoe",
		Password:    "password",
		PhoneNumber: 1234567890,
		Gender:      "Male",
	}

	user2 := models.User{
		Name:        "Jane Doe",
		Email:       "jane.doe@example.com",
		Username:    "janedoe",
		Password:    "password",
		PhoneNumber: 1234567891,
		Gender:      "Female",
	}

	// Create sample new health information
	user1.HealthInformation = models.HealthInformation{ // Data for User 1
		Height:            180,
		Weight:            70,
		BMI:               21.6,
		Birthday:          time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		ExerciseFrequency: 3,
	}

	user2.HealthInformation = models.HealthInformation{ // Data for User 2
		Height:            160,
		Weight:            80,
		BMI:               31.2,
		Birthday:          time.Date(1990, 2, 1, 0, 0, 0, 0, time.UTC),
		ExerciseFrequency: 3,
	}

	exerciseInfo1 := models.ExerciseInformation{
		ActivityName: "Running",
		Resistance:   "High",
	}

	exerciseInfo2 := models.ExerciseInformation{
		ActivityName: "Circuit training, minimal rest",
		Resistance:   "High",
	}

	exerciseInfo3 := models.ExerciseInformation{
		ActivityName: "Weight lifting, light workout",
		Resistance:   "Low",
	}

	// Add association for the users and their performed exercises
	user1.Exercises = []models.ExerciseInformation{exerciseInfo1, exerciseInfo2}
	user2.Exercises = []models.ExerciseInformation{exerciseInfo1, exerciseInfo3}

	// Save the information to the firestore database
	user1.Save()
	user2.Save()

	return models.FirestoreClient
}

// Setup the environment variables for the tests
// TOKEN_HOUR_LIFESPAN: 1
// API_SECRET: secret
func InitializeEnv(t *testing.T) {
	// Set the environment variables
	t.Setenv("TOKEN_HOUR_LIFESPAN", "1")
	t.Setenv("API_SECRET", "secret")
}
