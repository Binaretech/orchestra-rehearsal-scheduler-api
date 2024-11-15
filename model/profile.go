package model

type Profile struct {
	UserID    uint   `gorm:"type:bigint;primaryKey" json:"userId"`
	FirstName string `gorm:"type:varchar(255);not null" json:"firstName"`
	LastName  string `gorm:"type:varchar(255);not null" json:"lastName"`
	User      *User  `gorm:"foreignKey:UserID" json:"user"`
	Timestamps
}
