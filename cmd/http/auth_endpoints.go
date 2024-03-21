package http

import (
	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/app/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoadAuthEndpoints(router *gin.Engine, database *gorm.DB) {
	user := router.Group("/user")
	{
		user.GET("/ping", auth.PingGet())
	}
}
