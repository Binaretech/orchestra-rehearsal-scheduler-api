package model

import "time"

type Timestamps struct {
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

type BigId struct {
	ID int64 `gorm:"primaryKey;autoIncrement;type:bigint" json:"id"`
}
