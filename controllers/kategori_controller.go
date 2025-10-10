package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetKategori(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Daftar kategori"})
}

func CreateKategori(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Kategori berhasil dibuat"})
}
