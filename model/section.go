package model

// Section represents the section entity
type Section struct {
	BigId
	Name        string       `gorm:"type:varchar(255);not null" json:"name"`
	FamilyID    uint         `json:"familyId"`
	Family      *Family      `gorm:"foreignKey:FamilyID" json:"family"`
	Instruments []Instrument `gorm:"foreignKey:SectionID" json:"instruments,omitempty"`
}
