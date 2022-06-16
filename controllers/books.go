package controllers

import (
	"net/http"
	

	"github.com/gin-gonic/gin"
	"github.com/ifechigo/gin-quik/models"
)

type CreateWalletInput struct {
	Firstname  string `json:"firstname" binding:"required"`
	Lastname string `json:"lastname" binding:"required"`
}

type UpdatewalletInput struct {
	Amount float32 `json:"amount"`
}

// GET /wallets
// Find all wallets
func FindWallets(c *gin.Context) {
	var wallets []models.Wallet
	models.DB.Find(&wallets)

	c.JSON(http.StatusOK, gin.H{"data": wallets})
}

// GET /wallets/:id
// Find a wallet
func FindWallet(c *gin.Context) {
	// Get model if exist
	var wallet models.Wallet
	if err := models.DB.Where("id = ?", c.Param("id")).First(&wallet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": wallet})
}

// POST /wallets
// Create new wallet
func CreateWallet(c *gin.Context) {
	// Validate input
	var input CreateWalletInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check if Firstname exists

	// Create wallet
	wallet := models.Wallet{Firstname: input.Firstname, Lastname: input.Lastname, Amount: 0.00}
	models.DB.Create(&wallet)

	c.JSON(http.StatusOK, gin.H{"data": wallet})
}

// PATCH /wallets/:id
// Update a wallet
func UpdateWallet(c *gin.Context) {
	// Get model if exist
	var wallet models.Wallet
	if err := models.DB.Where("id = ?", c.Param("id")).First(&wallet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdatewalletInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&wallet).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": wallet})
}

// DELETE /wallets/:id
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
