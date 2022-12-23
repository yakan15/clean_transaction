package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title string `gorm:"not null"`
	Body  string `gorm:"not null"`
}
