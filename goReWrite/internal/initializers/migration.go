package initializers

import (
	"DevOps-Project/internal/models"
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var SQLiteDB *gorm.DB

func ConnectSQLite() {
	var err error
	dsn := os.Getenv("DATABASE_PATH")
	SQLiteDB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
}

func MigrateUsers() {
	var users []models.User

	if err := DB.Find(&users).Error; err != nil {
		fmt.Println("AutoMigrate Error: ", err)
	}

	for _, user := range users {
		fmt.Println("AutoMigrate: ", user)
		if err := DB.Create(&user).Error; err != nil {
			fmt.Println("Fejl ved indsættelse af data i PostgreSQL for bruger:", user.Username, err)
		}
	}
}

func MigratePages() {
	var pages []models.Page

	if err := DB.Find(&pages).Error; err != nil {
		fmt.Println("AutoMigrate Error: ", err)
	}

	for _, page := range pages {
		fmt.Println("AutoMigrate: ", page)
		if err := DB.Create(&page).Error; err != nil {
			fmt.Println("Fejl ved indsættelse af data i PostgreSQL for side:", page.Title, err)
		}
	}
}
