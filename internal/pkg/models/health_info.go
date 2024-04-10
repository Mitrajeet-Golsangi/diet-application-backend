package models

import (
	"time"
)

// Struct to store the health information of the User
type HealthInformation struct {
	Gender            string    `firestore:"gender,omitempty"`             // Gender of the user values may include, male, female or other (BMI calculated as per male standards for other)
	Birthday          time.Time `firestore:"birthday,omitempty"`           // Birthday of the user
	Weight            uint8     `firestore:"weight,omitempty"`             // Weight of the user in kilograms
	Height            uint8     `firestore:"height,omitempty"`             // Height of the user in centimeters
	BMI               float32   `firestore:"bmi,omitempty"`                // Body Mass Index of the user calculated from the weight and height
	ExerciseFrequency uint8     `firestore:"exercise_frequency,omitempty"` // Weekly exercise frequency of the user
}

// Update the health information of the user with the new provided information
func (h *HealthInformation) UpdateData(inp HealthInformation) {
	h.Gender = inp.Gender
	h.Birthday = inp.Birthday
	h.Weight = inp.Weight
	h.Height = inp.Height
	h.ExerciseFrequency = inp.ExerciseFrequency
	h.BMI = float32(h.Weight) / ((float32(h.Height) / 100) * (float32(h.Height) / 100))
}