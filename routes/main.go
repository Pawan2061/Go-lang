package routes

import (
	"example/my-go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/signup", controllers.Signup)
		userGroup.POST("/login")

	}
}
