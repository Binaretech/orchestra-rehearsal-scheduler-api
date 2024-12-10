package model

// Concert represents the concert entity
type Concert struct {
	BigId
	Title             string           `gorm:"type:varchar(255);not null" json:"title"`
	ConcertDate       string           `gorm:"not null" json:"concertDate"`
	ConcertDateStatus string           `gorm:"type:enum('tentative', 'definitive');not null" json:"concertDateStatus"`
	Location          string           `gorm:"type:varchar(255);not null" json:"location"`
	Description       string           `gorm:"type:varchar(255)" json:"description"`
	Rehearsals        []Rehearsal      `gorm:"foreignKey:ConcertID" json:"rehearsals"`
	ConcertSections   []ConcertSection `gorm:"foreignKey:ConcertID" json:"sections"`
	Timestamps
}

type ConcertSection struct {
	BigId
	ConcertID uint    `json:"concertId"`
	SectionID uint    `json:"sectionId"`
	Section   Section `gorm:"foreignKey:SectionID" json:"section"`
	Stands    []Stand `gorm:"foreignKey:ConcertSectionID" json:"stands"`
	Timestamps
}

type Stand struct {
	BigId
	StandNumber      uint            `json:"standNumber"`
	ConcertSectionID uint            `json:"concertSectionId"`
	ConcertSection   *ConcertSection `gorm:"foreignKey:ConcertSectionID" json:"concertSection,omitempty"`
	Users            []User          `gorm:"many2many:stand_users" json:"users"`
}
