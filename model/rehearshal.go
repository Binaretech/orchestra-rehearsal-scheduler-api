package model

import "time"

type Rehearsal struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	RehearsalDate time.Time `gorm:"not null" json:"rehearsalDate"`
	RehearsalTime time.Time `gorm:"not null" json:"rehearsalTime"`
	Location      string    `gorm:"type:varchar(255);not null" json:"location"`
	RehearsalType string    `gorm:"type:enum('general', 'regular');not null" json:"rehearsalType"`
	ConcertID     uint      `gorm:"not null" json:"concertId"`
	Concert       Concert   `gorm:"foreignKey:ConcertID" json:"concert"`
	Users         []User    `gorm:"many2many:rehearsal_users" json:"users"`
	Timestamps
}

// RehearsalUser represents the attendance of users in rehearsals
type RehearsalUser struct {
	RehearsalID uint `gorm:"primaryKey"`
	UserID      uint `gorm:"primaryKey"`
	Attendance  bool
}
