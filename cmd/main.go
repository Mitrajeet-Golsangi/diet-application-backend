package main

import (
	"log"
	"net/http"

	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/app/auth"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/helpers"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/models"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// load the environment variables from the .env file
	err := godotenv.Load()

	if err != nil {
		log.SetPrefix("API Server | ")
		log.Fatalln("| Error loading the .env file !")
	}
}

func main() {
	// Connect to the database
	models.InitDatabase()

	// Create a new Gin router
	r := gin.Default()

	// Display the sitemap of the API
	//? Note: Update this sitemap after making changes to the API
	//? endpoints further down the line in the application
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"/user": gin.H{
				"/register": "POST | Create a new user in the database",
				"/login": "POST | Log in an existing user in the database",
		},
		})
	})

	// The user management API endpoints
	user := r.Group("/user")
	{
		// User Registration Endpoints
		user.GET("/register", helpers.MethodNotAllowed("GET Method not Allowed !"))
		user.POST("/register", auth.RegisterPost())
		
		// User Login Endpoints
		user.GET("/login", helpers.MethodNotAllowed("GET Method not Allowed !"))
		user.POST("/login", auth.LoginPost())
	}

	// Protected Routes
	//? Note: The JWT token is required to access these routes
	//? The JWT token can be obtained by logging in using the /user/login endpoint
	//? The JWT token must be passed in the Authorization header as a Bearer token
	protected := r.Group("/api/v1/admin")
	{
		protected.Use(middlewares.JwtAuthMiddleware())
		protected.GET("/user", auth.CurrentUser())
	}

	// listen and serve on 0.0.0.0:8000
	r.Run()
}
