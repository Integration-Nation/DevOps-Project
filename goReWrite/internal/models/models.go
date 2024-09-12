package models

import "gorm.io/gorm"

type Page struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Content  string `gorm:"type:text" json:"content"`
	Language string `gorm:"type:varchar(5)" json:"language"`
}
