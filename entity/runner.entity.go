package entity

import "gorm.io/gorm"

type Runner struct {
	gorm.Model
	ID    int `gorm:"primarykey;autoincrement;"`
	Speed int
}