package firebaseapp

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var Instance *firebase.App
var Ctx context.Context

// Initialize the firebase application and store it in the Instance variable
func Initialize() {

	// Initialize the Firebase Admin SDK service account with private key
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_KEY_FILE"))

	Ctx = context.Background()
	var err error

	// Initialize the Firebase app
	Instance, err = firebase.NewApp(Ctx, &firebase.Config{ProjectID: "seal-the-meal"}, opt)
	if err != nil {
		log.SetPrefix("Firebase | ")
		log.Fatalf("error initializing app: %v\n", err)
	}
}
