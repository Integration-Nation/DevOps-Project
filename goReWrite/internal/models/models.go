package models

import "time"

type Page struct {
	Title       string    `gorm:"type:text" json:"title"`
	URL         string    `gorm:"type:text" json:"url"`
	Language    string    `gorm:"type:varchar(5)" json:"language"`
	LastUpdated time.Time `gorm:"type:timestamp" json:"last_updated"`
	Content     string    `gorm:"type:text" json:"content"`
}

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
