package models

type Page struct {
	ID       uint   `gorm:"primaryKey"`
	Content  string `gorm:"type:text" json:"content"`
	Language string `gorm:"type:varchar(5)" json:"language"`
}
