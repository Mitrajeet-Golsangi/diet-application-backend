package models

import "gorm.io/gorm"

// Struct to store the exercise information of the User
type ExerciseInformation struct {
	gorm.Model
	ActivityName string `firestore:"activity_name,omitempty"` // Name of the activity
	Resistance   string `firestore:"resistance,omitempty"` // The vigour for the exercise (valid values include High, Medium or Low)
}
