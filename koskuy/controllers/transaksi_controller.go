package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"koskuy/config"
	"koskuy/models"
)

// GET /transaksi
func GetTransaksi(c *gin.Context) {
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak terautentikasi"})
		return
	}

	var transaksi []models.Transaksi
	if err := config.DB.
		Joins("JOIN users ON users.id = transaksis.user_id").
		Joins("JOIN kategoris ON kategoris.id = transaksis.kategori_id").
		Where("users.email = ?", email).
		Preload("User").
		Preload("Kategori").
		Find(&transaksi).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Daftar transaksi kamu",
		"data":    transaksi,
	})
}


// POST /transaksi
func CreateTransaksi(c *gin.Context) {
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak terautentikasi"})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		return
	}

	var input models.Transaksi
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaksi := models.Transaksi{
		UserID:     user.ID,
		KategoriID: input.KategoriID,
		Jumlah:     input.Jumlah,
		Keterangan: input.Keterangan,
		Tanggal:    time.Now(),
	}

	if err := config.DB.Create(&transaksi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Transaksi berhasil dibuat",
		"transaksi": transaksi,
	})
}
