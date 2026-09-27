package repository

import (
	"PerpusAPI/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindByUsername(username string) (model.User, error) {
	var user model.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *UserRepository) FindByID(id uint) (model.User, error) {
	var user model.User
	err := r.DB.First(&user, id).Error
	return user, err
}

func (r *UserRepository) Create(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) Update(user *model.User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepository) Delete(id uint) error {
	return r.DB.Delete(&model.User{}, id).Error
}

func (r *UserRepository) FindAllAnggota(cari string) ([]model.User, error) {
	var users []model.User
	err := r.DB.Where("role = ? AND nama ILIKE ?", "user", "%"+cari+"%").Find(&users).Error
	return users, err
}
