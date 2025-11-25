package main

import (
	"log"

	"github.com/itssiddhant/taskmanager/internal/database"
	"github.com/itssiddhant/taskmanager/pkg/server"
)

func main() {
	db := database.Connect()

	r := server.New(db)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
