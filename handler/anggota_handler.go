package handler

import (
	"net/http"
	"strconv"

	"PerpusAPI/service"

	"github.com/gin-gonic/gin"
)

type AnggotaHandler struct {
	Service *service.AnggotaService
}

func NewAnggotaHandler(s *service.AnggotaService) *AnggotaHandler {
	return &AnggotaHandler{Service: s}
}

type updateAnggotaInput struct {
	Nama     string `json:"nama" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password"`
}

func (h *AnggotaHandler) GetAll(c *gin.Context) {
	cari := c.Query("cari")
	anggota, err := h.Service.GetAll(cari)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, anggota)
}

func (h *AnggotaHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var input updateAnggotaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.Service.Update(uint(id), input.Nama, input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *AnggotaHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	if err := h.Service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Anggota berhasil dihapus"})
}
