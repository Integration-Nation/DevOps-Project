package initializers

import (
	"DevOps-Project/internal/models"
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var SqliteDB *gorm.DB

func ConnectSqlite() {
	var err error
	dsn := os.Getenv("DATABASE_PATH")
	SqliteDB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
}

func MigrateUsers() {
	var users []models.User

	if err := SqliteDB.Find(&users).Error; err != nil {
		fmt.Println("AutoMigrate Error: ", err)
	}

	fmt.Println("AutoMigrate: ", len(users))

	for _, user := range users {
		fmt.Println("AutoMigrate: ", user.Username)
		if err := DB.Create(&user).Error; err != nil {
			fmt.Println("Fejl ved indsættelse af data i PostgreSQL for bruger:", user.Username, err)
		}
	}

	fmt.Println("users migrated")
}

func MigratePages() {
	var pages []models.Page

	if err := SqliteDB.Find(&pages).Error; err != nil {
		fmt.Println("AutoMigrate Error: ", err)
	}

	fmt.Println("AutoMigrate: ", len(pages))

	for _, page := range pages {
		fmt.Println("AutoMigrate: ", page.Title)
		if err := DB.Create(&page).Error; err != nil {
			fmt.Println("Fejl ved indsættelse af data i PostgreSQL for side:", page.Title, err)
		}
	}

	fmt.Println("pages migrated")
}
