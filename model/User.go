package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:64;not null"`
	Avatar   string `gorm:"size:255"`
	phone    string `gorm:"size:22"`
	Password string `gorm:"size:128;not null"`
}
