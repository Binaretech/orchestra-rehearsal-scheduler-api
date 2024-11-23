package model

// Family represents a instruments family entity
type Family struct {
	BigId
	Name     string    `gorm:"type:varchar(255);not null" json:"name"`
	Sections []Section `gorm:"foreignKey:SectionID" json:"instruments,omitempty"`
}
