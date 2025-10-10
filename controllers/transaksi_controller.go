package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetTransaksi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Daftar transaksi"})
}

func CreateTransaksi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Transaksi berhasil dibuat"})
}
