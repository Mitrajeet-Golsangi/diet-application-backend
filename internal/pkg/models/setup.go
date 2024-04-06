package models

import (
	"log"

	"cloud.google.com/go/firestore"
	firebaseapp "github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/firebase_app"
)

var FirestoreClient *firestore.Client
var UserCollection *firestore.CollectionRef

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
}
