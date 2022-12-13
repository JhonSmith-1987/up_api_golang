package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model

	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"not null;unique_index" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
