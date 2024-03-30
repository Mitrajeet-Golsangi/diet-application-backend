package models

import (
	"time"

	"gorm.io/gorm"
)

// Struct to store the health information of the User
type HealthInformation struct {
	gorm.Model
	UserID            uint      // User to which the health information belongs
	Birthday          time.Time // Birthday of the user
	Weight            uint8     // Weight of the user in kilograms
	Height            uint8     // Height of the user in centimeters
	BMI               float32   // Body Mass Index of the user calculated from the weight and height
	ExerciseFrequency uint8     // Weekly exercise frequency of the user
}
