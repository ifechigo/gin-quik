package controllers

import (
	"fmt"
	"net/http"

	"github.com/shopspring/decimal"

	"github.com/gin-gonic/gin"
	"github.com/ifechigo/gin-quik/models"
)

type CreateWalletInput struct {
	Firstname  string `json:"firstname" binding:"required"`
	Lastname string `json:"lastname" binding:"required"`
}

type UpdateWalletInput struct {
	Firstname  string `json:"firstname"`
	Lastname string `json:"lastname"`
}

type CreditWalletInput struct {
	Credit string `json:"amount"`
}

type DebitWalletInput struct {
	Debit string `json:"amount"`
}

// GET /api/v1/wallets
// Find all wallets
func FindWallets(c *gin.Context) {
	var wallets []models.Wallet

	models.DB.Find(&wallets)

	c.JSON(http.StatusOK, gin.H{"data": wallets})
}

// GET /api/v1/wallets/:id
// Find a wallet
func FindWallet(c *gin.Context) {
	// Get model if exist
	var wallet models.Wallet
	if err := models.DB.Where("id = ?", c.Param("id")).First(&wallet).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": wallet})
}

// POST /api/v1/wallets
// Create new wallet
func CreateWallet(c *gin.Context) {
	// Validate input
	var input CreateWalletInput
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

// PUT /api/v1/wallets/:id
// Update a wallet
func UpdateWallet(c *gin.Context) {
	// Get model if exist
	var wallet models.Wallet
	if err := models.DB.Where("id = ?", c.Param("id")).First(&wallet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateWalletInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&wallet).Updates(models.Wallet{
		Firstname: input.Firstname, Lastname: input.Lastname})

	c.JSON(http.StatusOK, gin.H{"data": wallet})
}

//PUT /api/v1/wallet/:id/credit
//Credit a wallet
func CreditWallet(c *gin.Context) {
	// Get model if exist
	var wallet models.Wallet
	if err := models.DB.Where("id = ?", c.Param("id")).First(&wallet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input CreditWalletInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//calculating new balance
	WalletBalance, err := decimal.NewFromString(input.Credit)
	if err != nil {
		panic(err)
	}

	amount := fmt.Sprintf("%v", wallet.Amount)
	
	credit, _ := decimal.NewFromString(amount)
	newAmount := WalletBalance.Add(credit)

	models.DB.Model(&wallet).Updates(models.Wallet{Amount: newAmount.InexactFloat64()})

	c.JSON(http.StatusOK, gin.H{"data": wallet})
}

//PUT /api/v1/wallet/:id/debit
//Debit a wallet
func DebitWallet(c *gin.Context) {
	// Get model if exist
	var wallet models.Wallet
	if err := models.DB.Where("id = ?", c.Param("id")).First(&wallet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input DebitWalletInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//calculating new balance
	amount := fmt.Sprintf("%v", wallet.Amount)

	WalletBalance, err := decimal.NewFromString(amount)
	if err != nil {
		panic(err)
	}

	debit, _ := decimal.NewFromString(input.Debit)
	newAmount := WalletBalance.Sub(debit)

	
	models.DB.Model(&wallet).Updates(models.Wallet{Amount: newAmount.InexactFloat64()})

	c.JSON(http.StatusOK, gin.H{"data": wallet})
}

// DELETE /api/v1/wallets/:id
// Delete a wallet
func DeleteWallet(c *gin.Context) {
	// Get model if exist
	var wallet models.Wallet
	if err := models.DB.Where("id = ?", c.Param("id")).First(&wallet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&wallet)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
