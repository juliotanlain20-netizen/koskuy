package models

import "time"

type Transaksi struct {
    ID         uint      `json:"id" gorm:"primaryKey"`
    UserID     uint      `json:"user_id"`
    KategoriID uint      `json:"kategori_id"`
    Jumlah     float64   `json:"jumlah"`
    Keterangan string    `json:"keterangan"`
    Tanggal    time.Time `json:"tanggal"`

    // Tambahkan relasi biar GORM tahu
    User     User     `json:"user" gorm:"foreignKey:UserID"`
    Kategori Kategori `json:"kategori" gorm:"foreignKey:KategoriID"`
}
