package model

import (
	"time"
)

// Concert represents the concert entity
type Concert struct {
	BigId
	Title             string      `gorm:"type:varchar(255);not null" json:"title"`
	ConcertDate       time.Time   `gorm:"not null" json:"concertDate"`
	ConcertDateStatus string      `gorm:"type:enum('tentative', 'definitive');not null" json:"concertDateStatus"`
	Location          string      `gorm:"type:varchar(255);not null" json:"location"`
	Description       string      `gorm:"type:varchar(255)" json:"description"`
	Rehearsals        []Rehearsal `gorm:"foreignKey:ConcertID" json:"rehearsals"`
	Timestamps
}
