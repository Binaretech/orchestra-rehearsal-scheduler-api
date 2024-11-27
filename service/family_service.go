package service

import (
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/model"
	"gorm.io/gorm"
)

type FamilyService struct {
	db *gorm.DB
}

func NewFamilyService(db *gorm.DB) *FamilyService {
	return &FamilyService{db: db}
}

func (s *FamilyService) GetAllData() ([]model.Family, error) {
	var families []model.Family

	if err := s.db.Preload("Sections").Find(&families).Error; err != nil {
		return nil, err
	}

	return families, nil
}
