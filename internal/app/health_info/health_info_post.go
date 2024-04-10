package health_info

import (
	"encoding/json"

	firebaseapp "github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/firebase_app"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/helpers"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/helpers/token"
	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HealthInfoPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the request body into the HealthInformation struct
		var inpData models.HealthInformation
		var healthInfoOut models.HealthInformation

		// Throw an error if the request body is invalid
		if err := c.ShouldBindJSON(&inpData); err != nil {
			helpers.BadRequest(c, err.Error())
			return
		}
		// Get the user ID from the token
		userID, err := token.GetUserID(c)

		// Throw an error if the user ID extraction fails
		if err != nil {
			helpers.InternalServerError(c, "Failed to extract user ID from token !")
			return
		}

		// Get the Health Information document associated with the current user in the database
		dsnap, err := models.UserCollection.Doc(userID).Get(firebaseapp.Ctx)

		// Throw an error if the health information creation fails
		if err != nil {
			helpers.InternalServerError(c, "Failed to fetch health information !")
			return
		}

		// Map the document reference to the HealthInformation struct
		dsnap.DataTo(&healthInfoOut)

		// Update the data in the database with the new input
		healthInfoOut.UpdateData(inpData)

		// Convert the health information struct to a JSON object
		healthInfoJSON, err := json.Marshal(healthInfoOut)

		// Throw an error if the JSON conversion fails
		if err != nil {
			helpers.InternalServerError(c, "Failed to convert health information to JSON !")
			return
		}
		
		logrus.Info("Health Information Updated Successfully !")

		c.JSON(200, gin.H{
			"message": "Health Information Updated Successfully !",
			"healthInfo": healthInfoJSON,
		})
	}
}