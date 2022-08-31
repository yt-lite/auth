package service

import (
	"github.com/yt-lite/auth/dto"
	"github.com/yt-lite/auth/models"
	"github.com/yt-lite/auth/repository"
	"github.com/yt-lite/libs/errs"
)

type AuthService struct {
	authRepo repository.AuthRepo
}

func NewAuthService(authRepo repository.AuthRepo) AuthService {
	return AuthService{authRepo}
}

func (as *AuthService) Register(req dto.RegisterAuthRequest) *errs.AppError {
	return as.authRepo.Create(&models.Auth{
		Email: req.Email,
	})
}

func (as *AuthService) Login(req dto.LoginAuthRequest) (*models.Auth, *errs.AppError) {
	return as.authRepo.FindByEmail(
		req.Email,
	)
}
