package model

type Profile struct {
	UserID    uint   `gorm:"type:bigint;primaryKey" json:"userId,omitempty"`
	FirstName string `gorm:"type:varchar(255);not null" json:"firstName,omitempty"`
	LastName  string `gorm:"type:varchar(255);not null" json:"lastName,omitempty"`
	User      *User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Timestamps
}
