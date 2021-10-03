package entity

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	ID   int `gorm:"primarykey;autoincrement;"`
	Name string
}