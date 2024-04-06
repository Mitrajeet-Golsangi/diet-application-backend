package models

import (
	"html"
	"strings"

	firebaseapp "github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/firebase_app"
	"golang.org/x/crypto/bcrypt"
)

// Struct to store the basic information about the User
// One to One relation with the Health Information of the User
// One to Many relation with the Exercises performed by the User
type User struct {
	Name              string                `firestore:"name,omitempty"`         // Full Name of the user
	Username          string                `firestore:"username,omitempty"`     // Username of the user used for logging in to the application
	Password          string                `firestore:"password,omitempty"`     // Password of the user used for logging in to the application
	Email             string                `firestore:"email,omitempty"`        // Email of the user, might be autofilled if user signs up using Google
	PhoneNumber       int64                 `firestore:"phone_number,omitempty"` // Phone number of the user
	Gender            string                `firestore:"gender,omitempty"`       // Gender of the user values may include, male, female or other (BMI calculated as per male standards for other)
	HealthInformation HealthInformation     // Health information of the user
	Exercises         []ExerciseInformation // Exercises performed by the user
}

// Save the user to the database
func (u *User) Save() (*User, error) {
	// Encrypt the password before saving it to the database

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return &User{}, err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	// Save the user to the database
	if _, err := UserCollection.NewDoc().Set(firebaseapp.Ctx, u); err != nil {
		return &User{}, err
	}

	return u, nil
}

// // Try to log in the user with the username and password provided
// func CheckByUsername(username string, password string) (string, error) {
// 	var err error

// 	u := User{}

// 	err = DB.Model(User{}).Where("username = ?", username).Take(&u).Error

// 	if err != nil {
// 		return "", err
// 	}

// 	err = VerifyPassword(password, u.Password)

// 	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
// 		return "", err
// 	}

// 	token, err := token.GenerateToken(u.ID)

// 	if err != nil {
// 		return "", err
// 	}

// 	return token, nil
// }

// // Find the user from the database with the provided ID
// func GetUserByID(uid uint) (User, error) {
// 	var u User

// 	if err := DB.First(&u, uid).Error; err != nil {
// 		return u, errors.New("User not Found")
// 	}

// 	u.PrepareGive()

// 	return u, nil
// }

// // Return the user without the password
// func (u *User) PrepareGive() {
// 	u.Password = ""
// }

// // Verify the password provided in the login request with the hashed password stored in the database
// func VerifyPassword(password, hashedPassword string) error {
// 	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// }
