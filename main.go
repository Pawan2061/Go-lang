package main

import (
	"example/my-go-rest-api/database"
	"example/my-go-rest-api/routes"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	router := gin.Default()
	fmt.Println("_________________________")
	database.SetupDatabase()
	fmt.Println("_________________________")

	router.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "working server",
		})
	})
	routes.UserRoutes(router)
	if error := router.Run(); error != nil {
		fmt.Print("theres some error in the router setup")
	}

}
