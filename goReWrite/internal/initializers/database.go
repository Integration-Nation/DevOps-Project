package initializers

import (
	"os"

	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := os.Getenv("DATABASE_PATH")
	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}
