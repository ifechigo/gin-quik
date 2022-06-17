package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ifechigo/gin-quik/controllers"
	"github.com/ifechigo/gin-quik/middleware"

)

func InitRouter() *gin.Engine {
	router := gin.Default()

	
	//logger
	router.Use(middleware.Logger())

	// Routes
	router.GET("/api/v1/wallets", controllers.FindWallets)
	router.POST("/api/v1/wallets", controllers.CreateWallet)
	router.GET("/api/v1/wallets/:id/balance", controllers.WalletBalance)
	router.POST("/api/v1/wallets/:id/credit", controllers.CreditWallet)
	router.POST("/api/v1/wallets/:id/debit", controllers.DebitWallet)

	return router
}