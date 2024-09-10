package models

import "gorm.io/gorm"

type Page struct {
	gorm.Model
	Content  string `json:"content"`
	Language string `json:"language"`
}