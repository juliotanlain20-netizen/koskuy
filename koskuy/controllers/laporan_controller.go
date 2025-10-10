package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"koskuy/config"
	"koskuy/models"
)

func GetLaporan(c *gin.Context) {
	var transaksi []models.Transaksi
	var total float64

	if err := config.DB.Find(&transaksi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, t := range transaksi {
		total += t.Jumlah
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Laporan keuangan",
		"total":   total,
		"data":    transaksi,
	})
}
