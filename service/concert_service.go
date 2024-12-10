package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

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

func (s *ConcertService) Create(title string, date string, location string, isDefinitive bool, rehearshalDays []string, distributions []ConcertDistribution) (*model.Concert, error) {
	concert := &model.Concert{
		Title:           title,
		ConcertDate:     date,
		Location:        location,
		Rehearsals:      []model.Rehearsal{},
		ConcertSections: []model.ConcertSection{},
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

		concert.ConcertSections = append(concert.ConcertSections, concertSection)
	}

	concertJSON, err := json.MarshalIndent(concert, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal concert to JSON: %w", err)
	}

	if err := ioutil.WriteFile("concert.json", concertJSON, 0644); err != nil {
		return nil, fmt.Errorf("failed to write concert JSON to file: %w", err)
	}

	if err := s.db.Create(concert).Error; err != nil {
		return nil, err
	}

	return concert, nil
}
