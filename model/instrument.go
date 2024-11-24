package model

// Instrument represents the instrument entity
type Instrument struct {
	BigId
	Name     string    `gorm:"type:varchar(255);not null" json:"name"`
	Sections []Section `gorm:"many2many:instrument_section" json:"sections"`
	Timestamps
}
