package handler

import (
	"net/http"
	"strconv"

	"PerpusAPI/service"

	"github.com/gin-gonic/gin"
)

type TransaksiHandler struct {
	Service *service.TransaksiService
}

func NewTransaksiHandler(s *service.TransaksiService) *TransaksiHandler {
	return &TransaksiHandler{Service: s}
}

func (h *TransaksiHandler) Pinjam(c *gin.Context) {
	userIDVal, _ := c.Get("user_id")
	userID := userIDVal.(uint)

	bukuID, err := strconv.Atoi(c.Param("id_buku"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID buku tidak valid"})
		return
	}

	transaksi, err := h.Service.PinjamBuku(userID, uint(bukuID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, transaksi)
}

func (h *TransaksiHandler) Kembalikan(c *gin.Context) {
	idTransaksi, err := strconv.Atoi(c.Param("id_transaksi"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID transaksi tidak valid"})
		return
	}

	if err := h.Service.KembalikanBuku(uint(idTransaksi)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil dikembalikan"})
}

func (h *TransaksiHandler) Riwayat(c *gin.Context) {
	userIDVal, _ := c.Get("user_id")
	userID := userIDVal.(uint)
	cari := c.Query("cari")

	list, err := h.Service.GetRiwayat(userID, cari)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *TransaksiHandler) GetAll(c *gin.Context) {
	cari := c.Query("cari")
	list, err := h.Service.GetAll(cari)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}
