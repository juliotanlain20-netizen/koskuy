package main

import (
	"koskuy/config"           // buat konek ke database
	"koskuy/routes"
)

func main() {
	// 1️⃣ Hubungkan ke database
	config.ConnectDatabase()

	// 2️⃣ Buat instance server
	r := routes.SetupRoutes()

	// 4️⃣ Jalankan server di port 8080
	r.Run(":8080")
}
