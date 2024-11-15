package model

// User represents the user entity
type User struct {
	BigId
	Email       string       `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password    string       `gorm:"type:varchar(255);not null" json:"pasword"`
	Role        string       `gorm:"type:varchar(20);not null" json:"role"`
	Instruments []Instrument `gorm:"many2many:user_instruments" json:"instruments"`
	Profile     Profile      `gorm:"foreignKey:UserID" json:"profile"`
	Timestamps
}

// VisibleUser represents the user entity without the password
type VisibleUser struct {
	ID          uint         `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName   string       `gorm:"type:varchar(255);not null" json:"firstName"`
	LastName    string       `gorm:"type:varchar(255);not null" json:"lastName"`
	Phone       string       `gorm:"type:varchar(20)" json:"phone"`
	Email       string       `gorm:"type:varchar(255);unique;not null" json:"email"`
	Role        string       `gorm:"type:enum('admin', 'coordinator', 'musician');not null" json:"role"`
	Instruments []Instrument `gorm:"many2many:user_instruments" json:"instruments"`
	Timestamps
}
