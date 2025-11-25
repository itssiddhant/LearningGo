package server

import (
	"github.com/gin-gonic/gin"
	"github.com/itssiddhant/taskmanager/pkg/routes"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	routes.RegisterAuthRoutes(router, db)
	routes.RegisterTaskRoutes(router, db)
	return router
}
