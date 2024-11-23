package service

import (
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type InstrumentService struct {
	db *gorm.DB
}

func NewInstrumentService(db *gorm.DB) *InstrumentService {
	return &InstrumentService{db: db}
}

func (s *InstrumentService) Create(name string, sectionId int64) (*model.Instrument, error) {
	section := model.Instrument{
		Name:      name,
		SectionID: sectionId,
	}

	if err := s.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}, {Name: "section_id"}},
		DoNothing: true,
	}).Create(&section).Error; err != nil {
		return nil, err
	}

	return &section, nil
}
