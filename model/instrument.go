package model

// Instrument represents the instrument entity
type Instrument struct {
	BigId
	Name      string  `gorm:"type:varchar(255);not null" json:"name"`
	SectionID int64   `gorm:"not null" json:"sectionId"`
	Section   Section `gorm:"foreignKey:SectionID" json:"section"`
	Users     []User  `gorm:"many2many:user_instruments" json:"users"`
	Timestamps
}

// UserInstrument represents the association between users and instruments
type UserInstrument struct {
	UserID       uint `gorm:"primaryKey"`
	InstrumentID uint `gorm:"primaryKey"`
}
