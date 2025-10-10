package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetLaporan(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Laporan keuangan"})
}
