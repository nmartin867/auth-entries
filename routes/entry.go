package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nmartin867/auth-logs/controllers"
)

func AuthEntryRoutes(r *gin.Engine) {
	userGroup := r.Group("/auth")
	{
		// Todo Routes
		userGroup.GET("", controllers.AuthEntryIndex) // Read all

	}
}
