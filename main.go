package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"user-service/config"
	"user-service/routes"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while loading .env file")
	}
	config.InitDB()
}

func main() {
	// Print to confirm env variables
	dbHost := os.Getenv("DB_HOST")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("✅ DB HOST: %s\n", dbHost)
	fmt.Printf("✅ Server running on port: %s\n", port)

	// Initialize Gin
	router := gin.Default()
	routes.RegisterUserRoutes(router)
	// Simple health check endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Add more routes here (e.g., user registration, campaign creation)
	// test push
	// Run server
	err := router.Run(":" + port)
	if err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
