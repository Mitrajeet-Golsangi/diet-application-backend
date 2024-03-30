package models

import (
	"html"
	"strings"

	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/helpers/token"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Struct to store the basic information about the User
// One to One relation with the Health Information of the User
// One to Many relation with the Exercises performed by the User
type User struct {
	gorm.Model
	Name              string                 `gorm:"not null"`        // Full Name of the user
	Username          string                 `gorm:"not null;unique"` // Username of the user used for logging in to the application
	Password          string                 `gorm:"not null"`        // Password of the user used for logging in to the application
	Email             string                 `gorm:"not null;unique"` // Email of the user, might be autofilled if user signs up using Google
	PhoneNumber       uint64                 `gorm:"not null;unique"` // Phone number of the user
	Gender            string                 `gorm:"not null"`        // Gender of the user values may include, male, female or other (BMI calculated as per male standards for other)
	HealthInformation HealthInformation      // Health information of the user
	Exercises         []*ExerciseInformation `gorm:"many2many:user_exercises;"` // Exercises performed by the user
}

// Save the user to the database
func (u *User) Save(DB *gorm.DB) (*User, error) {

	if err := DB.Create(&u).Error; err != nil {
		return &User{}, err
	}

	return u, nil
}

// Encrypt the password for the user before saving it to the database
func (u *User) BeforeSave(tx *gorm.DB) error {
	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}

// Try to log in the user with the username and password provided
func CheckByUsername(username string, password string, DB *gorm.DB) (string, error) {
	var err error

	u := User{}

	err = DB.Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

// Verify the password provided in the login request with the hashed password stored in the database
func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
