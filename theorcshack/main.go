package main

import (
	"log"
	"theorcshack/api/handlers"
	"theorcshack/api/routes"
	"theorcshack/config"
	"theorcshack/db"

	"github.com/gin-gonic/gin"
)

func main() {
	// Loading configuration
	config.LoadConfig()

	// Initialising database
	db.InitDatabase()

	// Initialising sentiment model
	if err := handlers.InitSentimentModel(); err != nil {
		log.Fatalf("Failed to initialize sentiment model: %v", err)
	}

	// Setting up Gin router
	router := gin.Default()

	// Initialising routes
	routes.InitRoutes(router)

	// Run the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
