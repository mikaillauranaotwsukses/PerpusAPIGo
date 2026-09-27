package service

import (
	"PerpusAPI/model"
	"PerpusAPI/repository"

	"golang.org/x/crypto/bcrypt"
)

type AnggotaService struct {
	Repo *repository.UserRepository
}

func NewAnggotaService(repo *repository.UserRepository) *AnggotaService {
	return &AnggotaService{Repo: repo}
}

func (s *AnggotaService) GetAll(cari string) ([]model.User, error) {
	return s.Repo.FindAllAnggota(cari)
}

func (s *AnggotaService) Update(id uint, nama, username, password string) (model.User, error) {
	existing, err := s.Repo.FindByID(id)
	if err != nil {
		return model.User{}, err
	}

	existing.Nama = nama
	existing.Username = username

	if password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return model.User{}, err
		}
		existing.Password = string(hashed)
	}

	err = s.Repo.Update(&existing)
	return existing, err
}

func (s *AnggotaService) Delete(id uint) error {
	return s.Repo.Delete(id)
}
