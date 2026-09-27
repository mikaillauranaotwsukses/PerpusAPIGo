package model

import "time"

type Transaksi struct {
	IDTransaksi    uint       `gorm:"primaryKey;column:id_transaksi" json:"id_transaksi"`
	IDUser         uint       `gorm:"column:id_user" json:"id_user"`
	IDBuku         uint       `gorm:"column:id_buku" json:"id_buku"`
	TanggalPinjam  *time.Time `gorm:"column:tanggal_pinjam" json:"tanggal_pinjam"`
	TanggalKembali *time.Time `gorm:"column:tanggal_kembali" json:"tanggal_kembali"`
	Status         string     `gorm:"column:status;default:dipinjam" json:"status"`

	User User `gorm:"foreignKey:IDUser;references:ID" json:"user,omitempty"`
	Buku Buku `gorm:"foreignKey:IDBuku;references:IDBuku" json:"buku,omitempty"`
}

func (Transaksi) TableName() string {
	return "transaksi"
}
