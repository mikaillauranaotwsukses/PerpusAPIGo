package main

import (
	"os"

	"PerpusAPI/config"
	"PerpusAPI/handler"
	"PerpusAPI/model"
	"PerpusAPI/repository"
	"PerpusAPI/routes"
	"PerpusAPI/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	config.ConnectDatabase()
	config.DB.AutoMigrate(&model.User{}, &model.Buku{}, &model.Transaksi{})

	userRepo := repository.NewUserRepository(config.DB)
	bukuRepo := repository.NewBukuRepository(config.DB)
	transaksiRepo := repository.NewTransaksiRepository(config.DB)

	authService := service.NewAuthService(userRepo)
	bukuService := service.NewBukuService(bukuRepo)
	anggotaService := service.NewAnggotaService(userRepo)
	transaksiService := service.NewTransaksiService(transaksiRepo, bukuRepo)

	h := routes.Handlers{
		Auth:      handler.NewAuthHandler(authService),
		Buku:      handler.NewBukuHandler(bukuService),
		Anggota:   handler.NewAnggotaHandler(anggotaService),
		Transaksi: handler.NewTransaksiHandler(transaksiService),
	}

	router := gin.Default()
	routes.SetupRoutes(router, h)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
