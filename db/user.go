package db

import (
	"database/sql"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

var DB *gorm.DB

func InitializeConnection() {
	// Open a database connection
	db, err := gorm.Open(postgres.Open("host=localhost port=5432 user=postgres dbname=postgres password=postgres"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database !")
	}
	
	// Migrate the schema
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("Failed to migrate schema !")
	}
	DB = db
}

func main() {
	db, err := gorm.Open(postgres.Open("host=localhost port=5432 user=postgres dbname=postgres password=postgres"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create
	db.Create(&User{Name: "Mitrajeet Golsangi", Age: 18})
}
