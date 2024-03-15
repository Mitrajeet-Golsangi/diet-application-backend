package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		log.SetPrefix("Database | ")
		log.Fatalln("| Error connecting to the database !")
	}

	return db
}
