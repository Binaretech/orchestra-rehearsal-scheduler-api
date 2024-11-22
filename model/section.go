package model

// Section represents the section entity
type Section struct {
	BigId
	Name        string       `gorm:"type:varchar(255);not null" json:"name"`
	Instruments []Instrument `gorm:"foreignKey:SectionID" json:"instruments,omitempty"`
}
