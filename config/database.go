package config

import (
	"fmt"
	"log"
	//"os"
 "koskuy/models" 
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Ganti sesuai koneksi PostgreSQL kamu
	dsn := "host=localhost user=postgres password=juliotanlain172006 dbname=koskuy_db port=5432 sslmode=disable"
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal konek ke database:", err)
	}

	fmt.Println("✅ Berhasil konek ke database!")

DB.AutoMigrate(&models.User{}, &models.Kategori{}, &models.Transaksi{})
fmt.Println("✅ Migrasi database berhasil!")

}
