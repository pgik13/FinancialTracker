package models

import "gorm.io/gorm"

type Budget struct {
	gorm.Model

	UserID   uint    `json:"userid" gorm:"not null"`
	Category string  `json:"category" gorm:"not null"`
	Amount   float64 `json:"amount" gorm:"not null"`
}
