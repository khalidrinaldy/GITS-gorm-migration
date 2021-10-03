package entity

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	ID    int `gorm:"primarykey;autoincrement;"`
	ModelName string
}