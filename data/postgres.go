package data

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	db, err := gorm.Open(postgres.Open(os.Getenv("con")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}
