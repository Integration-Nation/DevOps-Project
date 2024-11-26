package initializers

import (
	"DevOps-Project/internal/monitoring"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := os.Getenv("POSTGRES_DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}


func StartDBMonitoring() {
    go func() {
        for {
            monitoring.UpdateDBMetrics(DB) // DB er den globale databaseforbindelse
            time.Sleep(1 * time.Second)
        }
    }()
}