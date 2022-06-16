package main

import (
	"github.com/ifechigo/gin-quik/controllers"
	"github.com/ifechigo/gin-quik/models"
	"github.com/ifechigo/gin-quik/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//logger
	router.Use(middleware.Logger())

	// Connect to database
	models.ConnectDatabase()

	// Routes
	router.GET("/api/v1/wallets", controllers.FindWallets)
	router.POST("/api/v1/wallets", controllers.CreateWallet)
	router.GET("/api/v1/wallets/:id/balance", controllers.WalletBalance)
	router.POST("/api/v1/wallets/:id/credit", controllers.CreditWallet)
	router.POST("/api/v1/wallets/:id/debit", controllers.DebitWallet)
	
	// Run the server
	router.Run(":5005")
}
