package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var DB *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while loading .env file")
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	fmt.Printf("✅ Connected to MySQL database! \n")
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
