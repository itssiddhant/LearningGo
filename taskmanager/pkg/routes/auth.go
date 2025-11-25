package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/itssiddhant/taskmanager/pkg/handlers"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(r *gin.Engine, db *gorm.DB) {
	handler := handlers.NewAuthHandler(db)
	auth := r.Group("/auth")
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
	}
}
