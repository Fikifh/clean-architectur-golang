package database

import (
	"log"

	"incentrick-restful-api/repository/user_repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) (*gorm.DB, error) {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
	return Instance, dbError
}

func Migrate() {
	Instance.AutoMigrate(&user_repository.UserModel{})
}
