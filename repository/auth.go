package repository

import (
	"github.com/yt-lite/auth/models"
	"github.com/yt-lite/libs/errs"
	"gorm.io/gorm"
)

type AuthRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) AuthRepo {
	return AuthRepo{db}
}

func (ar *AuthRepo) Create(auth *models.Auth) *errs.AppError {
	if err := ar.db.Create(auth).Error; err != nil {
		return errs.NewInternalServerError(err.Error())
	}
	return nil
}

func (ar *AuthRepo) FindByID(id string) (*models.Auth, *errs.AppError) {
	auth := &models.Auth{}
	if err := ar.db.First(auth, id).Error; err != nil {
		return nil, errs.NewNotFoundError(err.Error())
	}
	return auth, nil
}

func (ar *AuthRepo) FindByEmail(email string) (*models.Auth, *errs.AppError) {
	auth := &models.Auth{}
	if err := ar.db.Where("email = ?", email).First(auth).Error; err != nil {
		return nil, errs.NewNotFoundError(err.Error())
	}
	return auth, nil
}
