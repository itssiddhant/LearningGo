package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/itssiddhant/taskmanager/internal/controllers"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(r *gin.Engine, db *gorm.DB) {
	authController := controllers.NewAuthController(db)
	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}
}
