package models

import (
	"time"
)

// Struct to store the health information of the User
type HealthInformation struct {
	Birthday          time.Time `firestore:"birthday,omitempty"`// Birthday of the user
	Weight            uint8     `firestore:"weight,omitempty"`// Weight of the user in kilograms
	Height            uint8     `firestore:"height,omitempty"`// Height of the user in centimeters
	BMI               float32   `firestore:"bmi,omitempty"`// Body Mass Index of the user calculated from the weight and height
	ExerciseFrequency uint8     `firestore:"exercise_frequency,omitempty"`// Weekly exercise frequency of the user
}
