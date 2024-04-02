package auth

import (
	"encoding/json"
	"log"

	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/helpers"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HealthInfoPost(DB *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the request body into the HealthInformation struct
		var healthInfo models.HealthInformation

		// Throw an error if the request body is invalid
		if err := c.ShouldBindJSON(&healthInfo); err != nil {
			helpers.BadRequest(c, err.Error())
			return
		}

		// Create a new health information in the database
		err := DB.Create(&healthInfo).Error

		// Throw an error if the health information creation fails
		if err != nil {
			helpers.InternalServerError(c, "Failed to create health information !")
			return
		}

		// Convert the health information struct to a JSON object
		healthInfoJSON, err := json.Marshal(healthInfo)

		// Throw an error if the JSON conversion fails
		if err != nil {
			helpers.InternalServerError(c, "Failed to convert health information to JSON !")
			return
		}

		log.Println("New Health Information Added Successfully !")

		c.JSON(200, gin.H{
			"message": "Health Information Added Successfully !",
			"healthInfo": healthInfoJSON,
		})
	}
}