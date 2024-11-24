package model

const (
	USER_ADMIN_ROLE       = "admin"
	USER_COORDINATOR_ROLE = "coordinator"
	USER_MUSICIAN_ROLE    = "musician"
)

// User represents the user entity
type User struct {
	BigId
	Email       string       `gorm:"type:varchar(255);unique;not null" json:"email,omitempty"`
	Password    string       `gorm:"type:varchar(255);not null" json:"-"`
	Role        string       `gorm:"type:varchar(20);not null" json:"role,omitempty"`
	Instruments []Instrument `gorm:"many2many:user_instruments" json:"instruments,omitempty"`
	Profile     Profile      `gorm:"foreignKey:UserID" json:"profile,omitempty"`
	Timestamps
}
