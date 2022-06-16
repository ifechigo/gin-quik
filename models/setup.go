package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DB_USERNAME = "root"
	DB_PASSWORD = "root"
	DB_NAME = "my_data"
	DB_HOST = "127.0.0.1"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := DB_USERNAME +":"+ DB_PASSWORD +"@tcp"+
			"(" + DB_HOST + ")/" + DB_NAME + "?" +
			"parseTime=true&loc=Local"  
	
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	database.AutoMigrate(&Wallet{})
	log.Println("Databasee Migration Completed!")

	DB = database
}


