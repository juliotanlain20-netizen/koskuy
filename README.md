Struktur & Inisialisasi Program 
Struktur Utama Kode:
main.go  untuk titik awal program, memanggil router
config/config.go untuk koneksi ke database PostgreSQL
models/ untuk definisi struct table (user, kategori, tranksasi)
controllers/ untuk logika CRUD dan autentikasi
middlewares/  pemeriksaan token JWT
routes/ untuk Menyusun rute API
main.go
func main() {
    config.ConnectDB()
    r := gin.Default()
    routes.UserRoutes(r)
    routes.KategoriRoutes(r)
    routes.TransaksiRoutes(r)
    r.Run(":8080")
}untuk menjalankan server di port 8080 dan menghubungkan semua route API

Models/ user.go

type User struct {
    gorm.Model
    Nama     string json:"nama"
    Email    string json:"email" gorm:"unique"
    Password string json:"password"
}
Kode ini untuk membuat struktur tabel user dengan validasi unik untuk email

Model & Database
Models/ user.go

type User struct {
    gorm.Model
    Nama     string json:"nama"
    Email    string json:"email" gorm:"unique"
    Password string json:"password"
}
Kode ini untuk membuat struktur tabel user dengan validasi unik untuk email
config/ config.go
dsn := "host=localhost user=postgres password=123 dbname=koskuy port=5432"
DB, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
DB.AutoMigrate(&models.User{}, &models.Kategori{}, &models.Transaksi{})
Untuk menghubungkan ke PostgreSQL dan membuat table otomatis
Modetype Transaksi struct {
    gorm.Model
    UserID     uint
    KategoriID uint
    Jumlah     float64
    Tanggal    time.Time
}
ls/ transaksi.go
Relasi antar table menggunakan foreign key (userID, KategoriID)

Controller & Middleware
controllers/ userController.go

func Register(c *gin.Context) {
    var user models.User
    c.BindJSON(&user)
    user.Password, _ = bcrypt.GenerateFromPassword([]byte(user.Password), 10)
    config.DB.Create(&user)
    c.utils/ jwt.go
func GenerateToken(id uint) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": id,
        "exp": time.Now().Add(time.Hour * 24).Unix(),
    })
    return token.SignedString([]byte("secret"))
}
Membuat toke JWT untuk autentikasi pengguna
JSON(200, gin.H{"message": "Registrasi berhasil"})
}
Kode ini menerima data JSON, mengenkripsi password, dan menyimpan ke database
middlewares/authMiddleware.go
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        // validasi token, tolak akses jika tidak sah
    }
}
Memeriksa token sebelum user bisa mengakses fitur lain (seperti tambah transaksi)
