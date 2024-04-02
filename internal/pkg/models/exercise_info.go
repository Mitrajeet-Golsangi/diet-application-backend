package models

import "gorm.io/gorm"

// Struct to store the exercise information of the User
type ExerciseInformation struct {
	gorm.Model
	ActivityName      string  // Name of the activity
	Category          string  // Type of activity (e.g. Cardio, Resistance, Sports, etc.)
	Resistance        string  // The vigour for the exercise (valid values include High, Medium or Low)
	EnergyExpenditure float32 // Energy expenditure of the activity in kilocalories per unit body weight of the person
	Users             []*User `gorm:"many2many:user_exercises;"`
}
