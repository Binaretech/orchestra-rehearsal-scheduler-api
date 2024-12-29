package service

import (
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/model"
	"gorm.io/gorm"
)

type MusicStand struct {
	Musicians []uint `json:"musicians" validate:"required,dive,required"`
	Stand     uint   `json:"stand" validate:"required"`
}

type ConcertDistribution struct {
	Section     uint         `json:"section" validate:"required"`
	MusicStands []MusicStand `json:"musicStands" validate:"required,dive,required"`
}

type ConcertService struct {
	db *gorm.DB
}

func NewConcertService(db *gorm.DB) *ConcertService {
	return &ConcertService{db: db}
}

func (s *ConcertService) Find(id uint) (*model.Concert, error) {
	var concert model.Concert

	if err := s.db.Preload("Rehearsals").Preload("Sections.Stands.Users.Profile").First(&concert, id).Error; err != nil {
		return nil, err
	}

	return &concert, nil
}

func (s *ConcertService) Create(title string, date string, location string, isDefinitive bool, rehearshalDays []string, distributions []ConcertDistribution) (*model.Concert, error) {
	concert := &model.Concert{
		Title:       title,
		ConcertDate: date,
		Location:    location,
		Rehearsals:  []model.Rehearsal{},
		Sections:    []model.ConcertSection{},
	}

	if isDefinitive {
		concert.ConcertDateStatus = "definitive"
	} else {
		concert.ConcertDateStatus = "tentative"
	}

	for _, rehearsalDay := range rehearshalDays {
		concert.Rehearsals = append(concert.Rehearsals, model.Rehearsal{Date: rehearsalDay})
	}

	for _, distribution := range distributions {
		concertSection := model.ConcertSection{SectionID: distribution.Section}

		for _, musicStand := range distribution.MusicStands {
			stand := model.Stand{StandNumber: musicStand.Stand}

			for _, musician := range musicStand.Musicians {
				stand.Users = append(stand.Users, model.User{
					BigId: model.BigId{
						ID: musician,
					},
				})
			}

			concertSection.Stands = append(concertSection.Stands, stand)
		}

		concert.Sections = append(concert.Sections, concertSection)
	}

	if err := s.db.Create(concert).Error; err != nil {
		return nil, err
	}

	concertData, err := s.Find(concert.ID)
	
	if err != nil {
		return nil, err
	}

	return concertData, nil
}
