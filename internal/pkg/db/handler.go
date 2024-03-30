package db

import (
	"log"
	"os"
	"time"

	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDatabase() *gorm.DB {

	// Configure database logging for the server
	consoleLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{
		Logger: consoleLogger,
	})

	if err != nil {
		log.SetPrefix("Database | ")
		log.Fatalln("| Error connecting to the database !")
		panic("Failed to connect database !")
	}

	// Create a list containing all the above structs
	// and pass it to the AutoMigrate function to create the tables in the database
	err = db.AutoMigrate(&models.User{}, &models.HealthInformation{}, &models.ExerciseInformation{})
	if err != nil {
		log.Fatal("Failed to migrate the schemas !")
		panic("Failed to migrate the schemas !")
	}

	return db
}
