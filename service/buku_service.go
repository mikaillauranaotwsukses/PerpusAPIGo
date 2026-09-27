package service

import (
	"PerpusAPI/model"
	"PerpusAPI/repository"
)

type BukuService struct {
	Repo *repository.BukuRepository
}

func NewBukuService(repo *repository.BukuRepository) *BukuService {
	return &BukuService{Repo: repo}
}

func (s *BukuService) GetAll(cari string) ([]model.Buku, error) {
	return s.Repo.FindAll(cari)
}

func (s *BukuService) GetByID(id uint) (model.Buku, error) {
	return s.Repo.FindByID(id)
}

func (s *BukuService) Create(buku *model.Buku) error {
	return s.Repo.Create(buku)
}

func (s *BukuService) Update(buku *model.Buku) error {
	return s.Repo.Update(buku)
}

func (s *BukuService) Delete(id uint) error {
	return s.Repo.Delete(id)
}
