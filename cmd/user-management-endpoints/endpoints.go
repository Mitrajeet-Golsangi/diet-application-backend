package usermanagementendpoints

import (
	usermanagementapi "github.com/Mitrajeet-Golsangi/diet-application-backend/internal/app/user-management-api"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoadEndpoints(router *gin.Engine, database *gorm.DB) {
	user := router.Group("/user")
	{
		user.GET("/ping", usermanagementapi.PingGet())
	}
}