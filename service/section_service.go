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

func (s *SectionService) GetAll() []*model.Section {
	sections := []*model.Section{}

	s.db.Find(&sections)

	return sections
}

func (s *SectionService) GetPaginated(page int, limit int) []*model.Section {
	sections := []*model.Section{}

	s.db.Offset((page - 1) * limit).Limit(limit).Find(&sections)

	return sections
}

func (s *SectionService) GetByID(id uint) *model.Section {
	section := &model.Section{}

	if err := s.db.Preload("Instruments").First(section, id).Error; err != nil {
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

func (s *SectionService) Create(name string, intrumentId int64) *model.Section {
	section := &model.Section{Name: name, InstrumentID: intrumentId}

	s.db.Create(section)

	return section
}

func (s *SectionService) Update(section *model.Section) *model.Section {
	s.db.Save(section)

	return section
}

func (s *SectionService) Delete(section *model.Section) {
	s.db.Delete(section)
}
