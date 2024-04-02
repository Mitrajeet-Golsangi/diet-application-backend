package testdata

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// Initialize the testing database with the required tables and sample data
// Users:
// 		John Doe
//		Jane Doe
// Health Information: 
// 		John Doe => Weight: 70, Height: 180, BMI: 21.6, Birthday: 1990-01-01, Exercise Frequency: 3
//		Jane Doe => Weight: 80, Height: 160, BMI: 31.2, Birthday: 1990-02-01, Exercise Frequency: 3
// Exercise Information:
// 		Running 
// 		Circuit training, minimal rest
//		Weight lifting, light workout
func InitializeTestDB() sqlmock.Sqlmock {
	mockDB, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	})

	DB, _ := gorm.Open(dialector, &gorm.Config{})

	// Assign the test database to the model database variable used globally
	models.DB = DB

	// Expect creation of the user table
	mock.ExpectQuery("SELECT .*").WillReturnRows(sqlmock.NewRows([]string{"version"}))
	mock.ExpectExec("CREATE TABLE \"users\" (.+)").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec("CREATE INDEX IF NOT EXISTS \".+\" ON \".+\" (.+)").WillReturnResult(sqlmock.NewResult(0, 0))
	
	// Expect creation of the exercise information table
	mock.ExpectQuery("SELECT .*").WillReturnRows(sqlmock.NewRows([]string{"version"}))
	mock.ExpectExec("CREATE TABLE \"exercise_informations\" (.+)").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec("CREATE INDEX IF NOT EXISTS \".+\" ON \".+\" (.+)").WillReturnResult(sqlmock.NewResult(0, 0))
	
	// Expect creation of the user_exercises table
	mock.ExpectQuery("SELECT .*").WillReturnRows(sqlmock.NewRows([]string{"version"}))
	mock.ExpectExec("CREATE TABLE \"user_exercises\" (.+)").WillReturnResult(sqlmock.NewResult(0, 0))
	
	// Expect creation of the health information table
	mock.ExpectQuery("SELECT .*").WillReturnRows(sqlmock.NewRows([]string{"version"}))
	mock.ExpectExec("CREATE TABLE \"health_informations\" (.+)").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec("CREATE INDEX IF NOT EXISTS \".+\" ON \".+\" (.+)").WillReturnResult(sqlmock.NewResult(0, 0))
	
	DB.AutoMigrate(&models.User{}, &models.HealthInformation{}, &models.ExerciseInformation{})

	// Create sample new users
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

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO \"users\" (.+) VALUES (.+)").WillReturnRows(mock.NewRows([]string{"id"}).AddRow("1"))
	mock.ExpectCommit()

	user1.BeforeSave(DB)
	u1, _ := user1.Save()

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO \"users\" (.+) VALUES (.+)").WillReturnRows(mock.NewRows([]string{"id"}).AddRow("1"))
	mock.ExpectCommit()

	user2.BeforeSave(DB)
	u2, _ := user2.Save()

	// Create sample new health information
	healthInfo1 := models.HealthInformation{ // Data for User 1
		Height:            180,
		Weight:            70,
		BMI:               21.6,
		Birthday:          time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		ExerciseFrequency: 3,
	}

	healthInfo2 := models.HealthInformation{ // Data for User 2
		Height:            160,
		Weight:            80,
		BMI:               31.2,
		Birthday:          time.Date(1990, 2, 1, 0, 0, 0, 0, time.UTC),
		ExerciseFrequency: 3,
	}

	DB.Create(&healthInfo1)
	DB.Create(&healthInfo2)

	// Associate the health information with the users
	DB.Model(u1).Association("HealthInformation").Append(&healthInfo1)
	DB.Model(u2).Association("HealthInformation").Append(&healthInfo2)

	// Create sample new exercise information
	exerciseInfo1 := models.ExerciseInformation{
		ActivityName:      "Running",
		Category:          "Cardio",
		Resistance:        "High",
		EnergyExpenditure: 1.75,
	}

	exerciseInfo2 := models.ExerciseInformation{
		ActivityName:      "Circuit training, minimal rest",
		Category:          "Resistance",
		Resistance:        "High",
		EnergyExpenditure: 1.64,
	}

	exerciseInfo3 := models.ExerciseInformation{
		ActivityName:      "Weight lifting, light workout",
		Category:          "Resistance",
		Resistance:        "Low",
		EnergyExpenditure: 0.61,
	}

	DB.Create(&exerciseInfo1)
	DB.Create(&exerciseInfo2)
	DB.Create(&exerciseInfo3)

	// Add association for the users and their performed exercises
	DB.Model(u1).Association("Exercises").Append([]models.ExerciseInformation{exerciseInfo1, exerciseInfo2})
	DB.Model(u2).Association("Exercises").Append([]models.ExerciseInformation{exerciseInfo1, exerciseInfo3})

	return mock
}

// Setup the environment variables for the tests
// TOKEN_HOUR_LIFESPAN: 1
// API_SECRET: secret
func InitializeEnv(t *testing.T) {
	// Set the environment variables
	t.Setenv("TOKEN_HOUR_LIFESPAN", "1")
	t.Setenv("API_SECRET", "secret")
}