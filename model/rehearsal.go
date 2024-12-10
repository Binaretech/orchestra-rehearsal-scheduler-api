package model

type Rehearsal struct {
	ID        uint     `gorm:"primaryKey;autoIncrement" json:"id"`
	Date      string   `gorm:"not null" json:"date"`
	Location  string   `gorm:"type:varchar(255);not null" json:"location"`
	IsGeneral bool     `gorm:"not null;default:false" json:"isGeneral"`
	ConcertID uint     `gorm:"not null" json:"concertId"`
	Concert   *Concert `gorm:"foreignKey:ConcertID" json:"concert"`
	Users     []User   `gorm:"many2many:rehearsal_users" json:"users"`
	Timestamps
}

// RehearsalUser represents the attendance of users in rehearsals
type RehearsalUser struct {
	RehearsalID uint `gorm:"primaryKey"`
	UserID      uint `gorm:"primaryKey"`
	Attendance  bool
}
