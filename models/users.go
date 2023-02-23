package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"not null; unique_index"`
	Task      []Task `gorm:"constraint:OnDelete:CASCADE"`
	Password  string `json:"password,omitempy"` //Le digo al json que no lo devuelva si está vacío
	Role      string `json:"role"`
}
