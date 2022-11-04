package db

import (
	"os"

	"github.com/juanvillacortac/entrenamiento-go/pkg/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: os.Getenv("DATABASE_URL")}))

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&entities.Song{})

	DB = db
}
