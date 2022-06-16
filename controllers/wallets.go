package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/shopspring/decimal"
	"github.com/gin-gonic/gin"

	"github.com/ifechigo/gin-quik/utils"
	"github.com/ifechigo/gin-quik/models"
	
)

// GET /api/v1/wallets
// Find all wallets
func FindWallets(c *gin.Context) {
	var wallets []models.Wallet

	models.DB.Find(&wallets)

	c.JSON(http.StatusOK, gin.H{"data": wallets})
}


// POST /api/v1/wallets
// Create new wallet
func CreateWallet(c *gin.Context) {
	// Validate input
	var input utils.CreateWalletInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create wallet
	wallet := models.Wallet{
		Firstname: input.Firstname, Lastname: input.Lastname, Amount: 0.00}

	models.DB.Create(&wallet)

	c.JSON(http.StatusOK, gin.H{"data": wallet})
}

//GET /api/v1/wallet/:id/balance
//Balance in Wallet
func WalletBalance(c *gin.Context) {
	// Get model if exist
	var wallet models.Wallet
	if err := models.DB.Where("id = ?", c.Param("id")).First(&wallet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input utils.CreditWalletInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Balance": wallet.Amount})
}



//POST /api/v1/wallet/:id/credit
//Credit a wallet
func CreditWallet(c *gin.Context) {
	// Get model if exist
	var wallet models.Wallet
	if err := models.DB.Where("id = ?", c.Param("id")).First(&wallet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input utils.CreditWalletInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//checking Negative Value
	amt, _ := strconv.ParseFloat(input.Credit, 32)
	if amt < 0.000 || input.Credit == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "amount cannot be Negative or Empty"})
		return
	}

	//calculating new balance
	walletBalance, err := decimal.NewFromString(input.Credit)
	if err != nil {
		panic(err)
	}

	amount := fmt.Sprintf("%v", wallet.Amount)
	credit, _ := decimal.NewFromString(amount)
	newAmount := walletBalance.Add(credit)

	models.DB.Model(&wallet).Updates(models.Wallet{Amount: newAmount.InexactFloat64()})

	c.JSON(http.StatusOK, gin.H{"data": wallet})
}

//POST /api/v1/wallet/:id/debit
//Debit a wallet
func DebitWallet(c *gin.Context) {
	// Get model if exist
	var wallet models.Wallet
	if err := models.DB.Where("id = ?", c.Param("id")).First(&wallet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input utils.DebitWalletInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	//checking Negative Value
	amt, _ := strconv.ParseFloat(input.Debit, 32)
	if amt < 0.000 || input.Debit == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "amount cannot be Negative or Empty"})
		return
	}

	//calculating new balance
	amount := fmt.Sprintf("%v", wallet.Amount)
	walletBalance, err := decimal.NewFromString(amount)
	if err != nil {
		panic(err)
	}

	debit, _ := decimal.NewFromString(input.Debit)
	newAmount := walletBalance.Sub(debit)

	if newAmount.InexactFloat64() < 0.00000000 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Balance cannot be less than zero"})
		return
	}
	
	models.DB.Model(&wallet).Updates(models.Wallet{Amount: newAmount.InexactFloat64()})

	c.JSON(http.StatusOK, gin.H{"data": wallet})
}


