package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model

	UserID   uint    `json:"userid" gorm:"not null"`
	Type     string  `json:"type" gorm:"not null"`     //incoming or outgoing
	Category string  `json:"category" gorm:"not null"` //labelling of type
	Amount   float64 `json:"amount" gorm:"not null"`
	Note     string  `json:"note"`
}
