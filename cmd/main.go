package main

import (
	"log"

	"github.com/Mitrajeet-Golsangi/diet-application-backend/cmd/http"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/db"
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
	database := db.InitDatabase()

	// Create a new Gin router
	r := gin.Default()

	// Display the sitemap of the API
	//? Note: Update this sitemap after making changes to the API
	//? endpoints further down the line in the application
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"User Management": "/user/ping | GET | Pings the user management API",
		})
	})

	// The user management API endpoints
	http.LoadAuthEndpoints(r, database)

	// listen and serve on 0.0.0.0:8080
	r.Run()
}
