package testdata

import (
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/db"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func InitializeTestDB() (*gorm.DB, sqlmock.Sqlmock) {
	mockDB, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	})
	
	DB, _ := gorm.Open(dialector, &gorm.Config{})

	DB.AutoMigrate(&db.User{}, &db.HealthInformation{}, &db.ExerciseInformation{})

	// Create sample new users
	user1 := db.User{
		Name:        "John Doe",
		Email:       "john.doe@example.com",
		Username:    "johndoe",
		Password:    "password",
		PhoneNumber: 1234567890,
		Gender:      "Male",
	}

	user2 := db.User{
		Name:        "Jane Doe",
		Email:       "jane.doe@example.com",
		Username:    "janedoe",
		Password:    "password",
		PhoneNumber: 1234567891,
		Gender:      "Female",
	}

	DB.Create(&user1)
	DB.Create(&user2)

	// Create sample new health information
	healthInfo1 := db.HealthInformation{ // Data for User 1
		Height:            180,
		Weight:            70,
		BMI:               21.6,
		Birthday:          time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		ExerciseFrequency: 3,
	}

	healthInfo2 := db.HealthInformation{ // Data for User 2
		Height:            160,
		Weight:            80,
		BMI:               31.2,
		Birthday:          time.Date(1990, 2, 1, 0, 0, 0, 0, time.UTC),
		ExerciseFrequency: 3,
	}
	
	DB.Create(&healthInfo1)
	DB.Create(&healthInfo2)

	// Associate the health information with the users
	DB.Model(&user1).Association("HealthInformation").Append(&healthInfo1)
	DB.Model(&user2).Association("HealthInformation").Append(&healthInfo2)

	// Create sample new exercise information
	exerciseInfo1 := db.ExerciseInformation{
		ActivityName:      "Running",
		Category:          "Cardio",
		Resistance:        "High",
		EnergyExpenditure: 1.75,
	}

	exerciseInfo2 := db.ExerciseInformation{
		ActivityName:      "Circuit training, minimal rest",
		Category:          "Resistance",
		Resistance:        "High",
		EnergyExpenditure: 1.64,
	}
	
	exerciseInfo3 := db.ExerciseInformation{
		ActivityName:      "Weight lifting, light workout",
		Category:          "Resistance",
		Resistance:        "Low",
		EnergyExpenditure: 0.61,
	}

	DB.Create(&exerciseInfo1)
	DB.Create(&exerciseInfo2)
	DB.Create(&exerciseInfo3)

	// Add association for the users and their performed exercises
	DB.Model(&user1).Association("Exercises").Append([]db.ExerciseInformation{exerciseInfo1, exerciseInfo2})
	DB.Model(&user2).Association("Exercises").Append([]db.ExerciseInformation{exerciseInfo1, exerciseInfo3})
	
	return DB, mock
}
