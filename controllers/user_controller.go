package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"koskuy/config"
	"koskuy/models"
)

// Ambil semua user
func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// Ambil user berdasarkan ID
func GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Update user
func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// Hapus user
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		return
	}

	config.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User berhasil dihapus"})
}
