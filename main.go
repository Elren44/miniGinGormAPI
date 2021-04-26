package main

import (
	"log"

	"github.com/Evelon44/MyTests/GormAPI/models"
	"github.com/Evelon44/MyTests/GormAPI/router"
	"github.com/Evelon44/MyTests/GormAPI/storage"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func main() {
	//connection string
	dsn := "host=localhost port=5432 user=postgres password=... dbname=myTest sslmode=disable"
	storage.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error while accessing database")
	}
	storage.DB.AutoMigrate(&models.ToDo{})

	r := router.RouterSetup()

	r.Run()
}
