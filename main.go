package main

import (
	"github.com/ifechigo/gin-quik/models"
	//"github.com/ifechigo/gin-quik/middleware"
	"github.com/ifechigo/gin-quik/router"
)

func main() {
	// Connect to database
	models.ConnectDatabase()

	router := router.InitRouter()

	//logger
	//router.Use(middleware.Logger())

	
	
	// Run the server
	router.Run(":5005")
}
