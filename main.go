package main

import (
	"github.com/ifechigo/gin-quik/controllers"
	"github.com/ifechigo/gin-quik/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	router.GET("/api/v1/wallets", controllers.FindWallets)
	router.GET("/api/v1/wallets/:id", controllers.FindWallet)
	router.POST("/api/v1/wallets", controllers.CreateWallet)
	router.PUT("/api/v1/wallets/:id", controllers.UpdateWallet)
	router.PUT("/api/v1/wallets/:id/credit", controllers.CreditWallet)
	router.PUT("/api/v1/wallets/:id/debit", controllers.DebitWallet)
	router.DELETE("/api/v1/wallets/:id", controllers.DeleteWallet)

	// Run the server
	router.Run()
}
