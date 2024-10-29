package routes

import "github.com/gin-gonic/gin"

func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/signup")
		userGroup.POST("/login")

	}
}
