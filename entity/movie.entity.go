package entity

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	ID    int `gorm:"primarykey;autoincrement;"`
	Title string
}