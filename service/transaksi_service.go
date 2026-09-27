package service

import (
	"errors"
	"time"

	"PerpusAPI/model"
	"PerpusAPI/repository"
)

type TransaksiService struct {
	TransaksiRepo *repository.TransaksiRepository
	BukuRepo      *repository.BukuRepository
}

func NewTransaksiService(tr *repository.TransaksiRepository, br *repository.BukuRepository) *TransaksiService {
	return &TransaksiService{TransaksiRepo: tr, BukuRepo: br}
}

func (s *TransaksiService) PinjamBuku(userID, bukuID uint) (model.Transaksi, error) {
	buku, err := s.BukuRepo.FindByID(bukuID)
	if err != nil {
		return model.Transaksi{}, errors.New("buku tidak ditemukan")
	}
	if buku.Stok <= 0 {
		return model.Transaksi{}, errors.New("stok buku habis")
	}

	now := time.Now()
	transaksi := model.Transaksi{
		IDUser:        userID,
		IDBuku:        bukuID,
		TanggalPinjam: &now,
		Status:        "dipinjam",
	}

	if err := s.TransaksiRepo.Create(&transaksi); err != nil {
		return model.Transaksi{}, err
	}

	if err := s.BukuRepo.UpdateStok(bukuID, -1); err != nil {
		return model.Transaksi{}, err
	}

	return transaksi, nil
}

func (s *TransaksiService) KembalikanBuku(idTransaksi uint) error {
	transaksi, err := s.TransaksiRepo.FindByID(idTransaksi)
	if err != nil {
		return errors.New("transaksi tidak ditemukan")
	}
	if transaksi.Status == "kembali" {
		return errors.New("buku sudah dikembalikan sebelumnya")
	}

	if err := s.TransaksiRepo.UpdateKembali(idTransaksi); err != nil {
		return err
	}

	return s.BukuRepo.UpdateStok(transaksi.IDBuku, 1)
}

func (s *TransaksiService) GetRiwayat(userID uint, cari string) ([]model.Transaksi, error) {
	return s.TransaksiRepo.FindRiwayatByUser(userID, cari)
}

func (s *TransaksiService) GetAll(cari string) ([]model.Transaksi, error) {
	return s.TransaksiRepo.FindAll(cari)
}
