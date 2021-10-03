package entity

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID   int `gorm:"primarykey;autoincrement;"`
	Title string
}