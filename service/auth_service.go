package service

import (
	"errors"

	"PerpusAPI/model"
	"PerpusAPI/repository"
	"PerpusAPI/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{UserRepo: repo}
}

func (s *AuthService) Register(nama, username, password string) (model.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		Nama:     nama,
		Username: username,
		Password: string(hashed),
		Role:     "user",
	}

	err = s.UserRepo.Create(&user)
	return user, err
}

func (s *AuthService) Login(username, password string) (string, model.User, error) {
	user, err := s.UserRepo.FindByUsername(username)
	if err != nil {
		return "", model.User{}, errors.New("username atau password salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", model.User{}, errors.New("username atau password salah")
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", model.User{}, err
	}

	return token, user, nil
}
