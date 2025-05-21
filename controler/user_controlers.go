package controler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"user-service/config"
	"user-service/models"
)

func RegisterUser(c *gin.Context) {

	//Bind the JSON Request
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("❌ Bind Error:", err) // Add this line
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check for existing User
	var existingUser models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	//hashing the Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}
	input.Password = string(hashedPassword)

	//Storing in DB
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "✅User created successfully"})
}
