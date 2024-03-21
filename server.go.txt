package main

import (
	"github.com/gin-gonic/gin"

	"diet.app/api/db"
)

func main() {
	r := gin.Default()

	// Initialize the database connection
	db.InitializeConnection()

	// API Functions
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, Diet Application !",
		})
	})

	// User CRUD Operations
	r.GET("/users", GetUsers)
	r.GET("/users/:id", GetUser)
	r.POST("/users", CreateUser)
	r.PUT("/users/:id", UpdateUser)
	r.DELETE("/users/:id", DeleteUser)

	// Run the server on 0.0.0.0:8080
	r.Run(":8080")
}

// Retrieves all the users from the database and returns them as a JSON response.
func GetUsers(c *gin.Context) {
	var users []db.User
	db.DB.Find(&users)

	c.JSON(200, users)
}

// Retrieves the user with the given id and returns it as a JSON response.
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user db.User
	db.DB.First(&user, id)

	c.JSON(200, user)
}

// Create a new user with the given
func CreateUser(c *gin.Context) {

	name := c.Param("name")

	// create a new user
	user := db.DB.Create(&db.User{Name: name, Age: 21})
	
	c.JSON(200, user)
}

// Update the user with the given id
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")

	// update the user
	user := db.DB.Model(&db.User{}).Where("id = ?", id).Update("name", name)
	
	c.JSON(200, user)
}

// Delete the user with the given id
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// delete the user
	user := db.DB.Delete(&db.User{}, id)
	
	c.JSON(200, user)
}