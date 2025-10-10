package routes

import (
	"github.com/gin-gonic/gin"
	"koskuy/controllers"
	"koskuy/middlewares"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to Koskuy API!"})
	})

	// Auth routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Protected routes (pakai JWT)
	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.GET("/users", controllers.GetUsers)
		auth.GET("/users/:id", controllers.GetUserByID)
		auth.PUT("/users/:id", controllers.UpdateUser)
		auth.DELETE("/users/:id", controllers.DeleteUser)

		auth.GET("/kategori", controllers.GetKategori)
		auth.POST("/kategori", controllers.CreateKategori)
		auth.GET("/transaksi", controllers.GetTransaksi)
		auth.POST("/transaksi", controllers.CreateTransaksi)
		auth.GET("/laporan", controllers.GetLaporan)
	}

	return r
}
