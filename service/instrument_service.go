package service

import (
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/model"
	"gorm.io/gorm"
)

type InstrumentService struct {
	db *gorm.DB
}

func NewInstrumentService(db *gorm.DB) *InstrumentService {
	return &InstrumentService{db: db}
}

func (s *InstrumentService) Create(name string) (*model.Instrument, error) {
	instrument := model.Instrument{
		Name: name,
	}

	if err := s.db.Create(&instrument).Error; err != nil {
		return nil, err
	}

	return &instrument, nil
}
