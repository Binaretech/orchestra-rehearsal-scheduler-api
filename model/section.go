package model

// Section represents the section entity
type Section struct {
	SectionID   uint         `gorm:"primaryKey;autoIncrement"`
	Name        string       `gorm:"type:varchar(255);not null"`
	Instruments []Instrument `gorm:"foreignKey:SectionID"`
}
