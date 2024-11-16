package service

import (
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/cache"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/model"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB, cache cache.Cache) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) GetByEmail(email string) (*model.User, error) {
	var user model.User

	if err := s.db.Joins("Profile").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
