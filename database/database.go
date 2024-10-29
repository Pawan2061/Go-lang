package database

import (
	"example/my-go-rest-api/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func SetupDatabase() {
	var err error

	dsn := "host=ep-icy-fire-a52bon09.us-east-2.aws.neon.tech user=Pawan2061 password=eTuqbHO0GJD8 dbname=go port=5432 sslmode=require"

	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = Db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
