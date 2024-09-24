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
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type Weather struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationTimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimeZoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	Current              Current `json:"current"`
	Hourly               Hourly  `json:"hourly"`
}

// Current represents the current weather data.
type Current struct {
	Time          string  `json:"time"`
	Interval      int     `json:"interval"`
	Temperature2m float64 `json:"temperature_2m"`
}

// Hourly represents the hourly weather data.
type Hourly struct {
	Time          []string  `json:"time"`
	Temperature2m []float64 `json:"temperature_2m"`
}
