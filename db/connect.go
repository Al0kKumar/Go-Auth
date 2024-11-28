package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	url := "postgres:localhost:5713"
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to postgres db %v", err)
	}

	log.Println("Database connected sucessfully")
	DB = db
}
