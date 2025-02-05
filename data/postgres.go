package data

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
  
var con = "host=localhost user=postgres password=admin dbname=UMC port=5432 sslmode=disable TimeZone=America/Lima"
var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open(con), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}