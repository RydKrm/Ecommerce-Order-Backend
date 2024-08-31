package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection(){
	db_url := os.Getenv("DB_URL");
	if db_url == ""{
		log.Fatal("Database url cannot found")
	}
    // db_url := "host=localhost user=postgres password=yourpassword dbname=yourdb port=5432 sslmode=disable"
    var err error
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{});
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB, err = gorm.Open(postgres.Open(db_url), &gorm.Config{})
	if err!= nil {
		log.Fatal("Failed to connect to database", err)
	}

	fmt.Println("Database connected")

}
