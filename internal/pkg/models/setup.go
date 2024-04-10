package models

import (
	"log"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/v4/auth"
	firebaseapp "github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/firebase_app"
)

var FirestoreClient *firestore.Client
var UserCollection *firestore.CollectionRef
var AuthClient *auth.Client

func InitDatabase() {
	var err error
	firebaseapp.Initialize()
	// Initialize the Firestore client
	FirestoreClient, err = firebaseapp.Instance.Firestore(firebaseapp.Ctx)
	if err != nil {
		log.SetPrefix("Firestore | ")
		log.Fatalf("error getting Firestore client: %v\n", err)
	}

	// Create the user information collection in the Firestore database
	UserCollection = FirestoreClient.Collection("user_info")

	// Initialize the Firebase Auth client
	AuthClient, err = firebaseapp.Instance.Auth(firebaseapp.Ctx)
	if err != nil {
		log.SetPrefix("Firebase Auth | ")
		log.Fatalf("error getting Firebase Auth client: %v\n", err)
	}
}
