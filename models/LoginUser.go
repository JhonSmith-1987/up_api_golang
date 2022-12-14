package models

import "gorm.io/gorm"

type LoginUser struct {
	gorm.Model

	Email    string `gorm:"not null;unique_index" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
