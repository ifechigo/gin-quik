package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

const (
	DB_USERNAME = "root"
	DB_PASSWORD = "root"
	DB_NAME = "my_data"
	DB_HOST = "127.0.0.1"
	DB_PORT = "3306"
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

	DB = database
}
