package repository

import (
	"time"

	"PerpusAPI/model"

	"gorm.io/gorm"
)

type TransaksiRepository struct {
	DB *gorm.DB
}

func NewTransaksiRepository(db *gorm.DB) *TransaksiRepository {
	return &TransaksiRepository{DB: db}
}

func (r *TransaksiRepository) Create(t *model.Transaksi) error {
	return r.DB.Create(t).Error
}

func (r *TransaksiRepository) FindByID(id uint) (model.Transaksi, error) {
	var t model.Transaksi
	err := r.DB.First(&t, id).Error
	return t, err
}

func (r *TransaksiRepository) UpdateKembali(id uint) error {
	now := time.Now()
	return r.DB.Model(&model.Transaksi{}).Where("id_transaksi = ?", id).
		Updates(map[string]interface{}{
			"tanggal_kembali": now,
			"status":          "kembali",
		}).Error
}

func (r *TransaksiRepository) FindRiwayatByUser(userID uint, cari string) ([]model.Transaksi, error) {
	var list []model.Transaksi
	err := r.DB.Preload("Buku").
		Where("id_user = ?", userID).
		Where("id_buku IN (?)", r.DB.Model(&model.Buku{}).Select("id_buku").Where("judul ILIKE ?", "%"+cari+"%")).
		Order("id_transaksi desc").
		Find(&list).Error
	return list, err
}

func (r *TransaksiRepository) FindAll(cari string) ([]model.Transaksi, error) {
	var list []model.Transaksi
	err := r.DB.Preload("Buku").Preload("User").
		Where("id_buku IN (?) OR id_user IN (?)",
			r.DB.Model(&model.Buku{}).Select("id_buku").Where("judul ILIKE ?", "%"+cari+"%"),
			r.DB.Model(&model.User{}).Select("id").Where("nama ILIKE ?", "%"+cari+"%"),
		).
		Order("id_transaksi desc").
		Find(&list).Error
	return list, err
}
