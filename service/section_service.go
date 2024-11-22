package service

import (
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/model"
	"gorm.io/gorm"
)

type SectionService struct {
	db *gorm.DB
}

func NewSectionService(db *gorm.DB) *SectionService {
	return &SectionService{db: db}
}

func (s *SectionService) GetByID(id uint) *model.Section {
	section := &model.Section{}

	if err := s.db.First(section, id).Error; err != nil {
		return nil
	}

	return section
}

func (s *SectionService) GetByName(name string) *model.Section {
	section := &model.Section{}

	if err := s.db.Where("name = ?", name).First(section).Error; err != nil {
		return nil
	}

	return section
}

func (s *SectionService) Create(name string) *model.Section {
	section := &model.Section{Name: name}

	s.db.Create(section)

	return section
}
