package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/itssiddhant/taskmanager/pkg/handlers"
	"github.com/itssiddhant/taskmanager/pkg/middleware"
	"gorm.io/gorm"
)

func RegisterTaskRoutes(r *gin.Engine, db *gorm.DB) {
	handler := handlers.NewTaskHandler(db)
	tasks := r.Group("/tasks")
	tasks.Use(middleware.AuthMiddleware())
	{
		tasks.POST("", handler.CreateTask)
		tasks.GET("", handler.GetTasks)
		tasks.PUT("/:id", handler.UpdateTask)
		tasks.DELETE("/:id", handler.DeleteTask)
	}
}
