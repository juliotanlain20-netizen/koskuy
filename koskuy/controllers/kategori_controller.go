package controllers

import (
	"net/http"
	"koskuy/config"
	"koskuy/models"

	"github.com/gin-gonic/gin"
)

// GET /kategori
func GetKategori(c *gin.Context) {
	var kategori []models.Kategori

	if err := config.DB.Find(&kategori).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Daftar kategori",
		"data":    kategori,
	})
}

// POST /kategori
func CreateKategori(c *gin.Context) {
	var input models.Kategori

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	kategori := models.Kategori{
		Nama: input.Nama,
	}

	if err := config.DB.Create(&kategori).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Kategori berhasil dibuat",
		"kategori": kategori,
	})
}
