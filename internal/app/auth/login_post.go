package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, nil)

		// var input LoginInput

		// // Bind the input data and validate the request body
		// if err := c.ShouldBindJSON(&input); err != nil {
		// 	helpers.BadRequest(c, err.Error())
		// 	return
		// }

		// // Check if the user exists in the database
		// token, err := models.CheckByUsername(input.Username, input.Password)

		// // Throw and error if the login fails
		// if err != nil {
		// 	helpers.Unauthorized(c, "Failed to login user !")
		// 	return
		// }

		// c.JSON(http.StatusOK, gin.H{"token": token})
	}
}