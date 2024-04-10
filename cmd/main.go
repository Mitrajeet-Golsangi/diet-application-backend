package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/app/auth"
	firebaseapp "github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/firebase_app"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/helpers"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/logging"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/models"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//? Note: In debug mode the application will try to use Firebase emulators
//! If not found the environment variables it will use the Production !
func init() {
	// Initialize the logging
	logging.InitializeLogger()

	log.SetPrefix("API Server | ")
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("| Error getting the current working directory !")
	}
	
	// load the environment variables from the .env file
	err = godotenv.Load(cwd + "/.env")
	
	log.SetPrefix("Firebase Admin SDK | ")
	log.Println("| GCLOUD_PROJECT: ", os.Getenv("GCLOUD_PROJECT"))
	log.Println("| FIREBASE_AUTH_EMULATOR_HOST: ", os.Getenv("FIREBASE_AUTH_EMULATOR_HOST"))
	log.Println("| FIRESTORE_EMULATOR_HOST: ", os.Getenv("FIRESTORE_EMULATOR_HOST"))

	log.SetPrefix("API Server | ")
	if err != nil {
		log.Println("| Error loading the .env file, working in Production mode !")
	}
	
	// Initialize the Firebase Application
	firebaseapp.Initialize()
	
	// Connect to the database
	models.InitDatabase()
}

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Set the base path for the API
	api := r.Group("/api/v1")

	// Display the sitemap of the API
	//? Note: Update this sitemap after making changes to the API
	//? endpoints further down the line in the application
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"/user": gin.H{
				"/register": "POST | Create a new user in the database",
				"/login":    "POST | Log in an existing user in the database",
			},
		})
	})

	// The user management API endpoints
	user := api.Group("/user")
	{
		// User Registration Endpoints
		user.GET("/register", helpers.MethodNotAllowed("GET Method not Allowed !"))
		user.POST("/register", auth.RegisterPost())

		//* Note: Login handled in the frontend by Firebase API
	}

	// Protected Routes
	//? Note: The JWT token is required to access these routes
	//? The JWT token can be obtained by logging in using the /user/login endpoint
	//? The JWT token must be passed in the Authorization header as a Bearer token
	api.Use(middlewares.JwtAuthMiddleware())
	
	// Get the user information
	api.GET("/user", auth.CurrentUser())

	// listen and serve on 0.0.0.0:8000
	r.Run()
	
	// Close the log file after the server stops
	logging.CloseLogger()
}
