package db

import (
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Dbconnection() {
	var error error
	DNS := os.Getenv("DSN")

	DB, error = gorm.Open(sqlite.Open(DNS), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Database connected successfully")
	}
}
