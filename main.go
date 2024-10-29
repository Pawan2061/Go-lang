package main

import (
	"example/my-go-rest-api/models"
	"example/my-go-rest-api/routes"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func setupDatabase() {
	db, err := gorm.Open("postgres", "postgresql://Pawan2061:eTuqbHO0GJD8@ep-icy-fire-a52bon09.us-east-2.aws.neon.tech/go?sslmode=require")
	if err != nil {
		panic("failed to connecrt")

	}
	db.AutoMigrate(&models.User{})

}

func main() {

	router := gin.Default()

	setupDatabase()

	router.GET("/", func(ctx *gin.Context) {
		fmt.Println("this is working")
	})
	routes.UserRoutes(router)
	if error := router.Run(); error != nil {
		fmt.Print("theres some error in the router setup")
	}

}
