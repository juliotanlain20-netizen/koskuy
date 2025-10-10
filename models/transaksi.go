package models

import "time"

type Transaksi struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `json:"user_id"`
	KategoriID  uint      `json:"kategori_id"`
	Nominal     float64   `json:"nominal"`
	Tanggal     time.Time `json:"tanggal"`
	Keterangan  string    `json:"keterangan"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
