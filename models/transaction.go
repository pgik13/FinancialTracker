package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model

	Type     string  `json:"type" gorm:"not null"`     //incoming or outgoing
	Category string  `json:"category" gorm:"not null"` //labelling of type
	Amount   float64 `json:"amount" gorm:"not null"`
	Note     string  `json:"note"`
	UserID   uint    `json:"userid" gorm:"not null"`
	User     User    `gorm:"foreignKey:UserID"`
}
