package database

import (
	"go-crud-api/entities"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect(ConnectionString string) {
	Instance, err = gorm.Open(mysql.Open(ConnectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to database")
	}
	log.Println("Connecting to database")
}

func Migrate() {
	Instance.AutoMigrate(&entities.Product{})
	log.Println("Database migration completed")
}
