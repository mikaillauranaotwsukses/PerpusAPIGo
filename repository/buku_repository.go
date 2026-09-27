package repository

import (
	"PerpusAPI/model"

	"gorm.io/gorm"
)

type BukuRepository struct {
	DB *gorm.DB
}

func NewBukuRepository(db *gorm.DB) *BukuRepository {
	return &BukuRepository{DB: db}
}

func (r *BukuRepository) FindAll(cari string) ([]model.Buku, error) {
	var buku []model.Buku
	err := r.DB.Where("judul ILIKE ?", "%"+cari+"%").Find(&buku).Error
	return buku, err
}

func (r *BukuRepository) FindByID(id uint) (model.Buku, error) {
	var buku model.Buku
	err := r.DB.First(&buku, id).Error
	return buku, err
}

func (r *BukuRepository) Create(buku *model.Buku) error {
	return r.DB.Create(buku).Error
}

func (r *BukuRepository) Update(buku *model.Buku) error {
	return r.DB.Save(buku).Error
}

func (r *BukuRepository) Delete(id uint) error {
	return r.DB.Delete(&model.Buku{}, id).Error
}

func (r *BukuRepository) UpdateStok(id uint, delta int) error {
	return r.DB.Model(&model.Buku{}).Where("id_buku = ?", id).
		UpdateColumn("stok", gorm.Expr("stok + ?", delta)).Error
}
