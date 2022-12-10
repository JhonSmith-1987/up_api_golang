package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model

	Name     string `gorm:"not null"`
	Email    string `gorm:"not null;unique_index"`
	Password string `gorm:"not null"`
}
