package routes

import (
	"PerpusAPI/handler"
	"PerpusAPI/middleware"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Auth      *handler.AuthHandler
	Buku      *handler.BukuHandler
	Anggota   *handler.AnggotaHandler
	Transaksi *handler.TransaksiHandler
}

func SetupRoutes(router *gin.Engine, h Handlers) {
	api := router.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/register", h.Auth.Register)
		auth.POST("/login", h.Auth.Login)
	}

	buku := api.Group("/buku")
	buku.Use(middleware.AuthRequired())
	{
		buku.GET("", h.Buku.GetAll)
		buku.GET("/:id", h.Buku.GetByID)
		buku.POST("", middleware.AdminOnly(), h.Buku.Create)
		buku.PUT("/:id", middleware.AdminOnly(), h.Buku.Update)
		buku.DELETE("/:id", middleware.AdminOnly(), h.Buku.Delete)
	}

	anggota := api.Group("/anggota")
	anggota.Use(middleware.AuthRequired(), middleware.AdminOnly())
	{
		anggota.GET("", h.Anggota.GetAll)
		anggota.PUT("/:id", h.Anggota.Update)
		anggota.DELETE("/:id", h.Anggota.Delete)
	}

	transaksi := api.Group("/transaksi")
	transaksi.Use(middleware.AuthRequired())
	{
		transaksi.POST("/pinjam/:id_buku", h.Transaksi.Pinjam)
		transaksi.POST("/kembali/:id_transaksi", h.Transaksi.Kembalikan)
		transaksi.GET("/riwayat", h.Transaksi.Riwayat)
		transaksi.GET("", middleware.AdminOnly(), h.Transaksi.GetAll)
	}
}
