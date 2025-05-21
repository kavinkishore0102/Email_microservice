package routes

import (
	"github.com/gin-gonic/gin"
	"user-service/controler"
)

func RegisterUserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", controler.RegisterUser)
	}
}
