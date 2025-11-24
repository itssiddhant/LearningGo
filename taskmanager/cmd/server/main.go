package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itssiddhant/taskmanager/internal/database"
	"github.com/itssiddhant/taskmanager/internal/routes"
)

func main() {
	db := database.Connect()
	r := gin.Default()
	routes.RegisterAuthRoutes(r, db)
	routes.RegisterTaskRoutes(r, db)
	r.Run(":8080")
}
