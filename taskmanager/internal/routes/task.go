package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/itssiddhant/taskmanager/internal/controllers"
	"github.com/itssiddhant/taskmanager/internal/middleware"
	"gorm.io/gorm"
)

func RegisterTaskRoutes(r *gin.Engine, db *gorm.DB) {
	taskController := controllers.NewTaskController(db)
	tasks := r.Group("/tasks")
	tasks.Use(middleware.AuthMiddleware())
	{
		tasks.POST("", taskController.CreateTask)
		tasks.GET("", taskController.GetTasks)
		tasks.PUT("/:id", taskController.UpdateTask)
		tasks.DELETE("/:id", taskController.DeleteTask)
	}
}
