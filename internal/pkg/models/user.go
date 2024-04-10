package models

import (
	firebaseapp "github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/firebase_app"
)

// Struct to store the basic information about the User
// One to One relation with the Health Information of the User
// One to Many relation with the Exercises performed by the User
type UserInfo struct {
	HealthInformation HealthInformation     // Health information of the user
	Exercises         []ExerciseInformation // Exercises performed by the user
}

// Save the user to the database
func (u *UserInfo) Save() (*UserInfo, error) {
	// Save the user information to the database
	if _, err := UserCollection.NewDoc().Set(firebaseapp.Ctx, u); err != nil {
		return &UserInfo{}, err
	}

	return u, nil
}

// Get the user information from the database
func (u *UserInfo) Get(userID string) (*UserInfo, error) {
	// Get the user information from the database
	doc, err := UserCollection.Doc(userID).Get(firebaseapp.Ctx)
	if err != nil {
		return &UserInfo{}, err
	}

	// Unmarshal the document into the user struct
	if err := doc.DataTo(u); err != nil {
		return &UserInfo{}, err
	}

	return u, nil
}