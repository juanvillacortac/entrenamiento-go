package db

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/juanvillacortac/entrenamiento-go/pkg/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RDB *redis.Client
var RCtx = context.Background()

func ConnectDatabase() {
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: os.Getenv("DATABASE_URL")}))

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&entities.Song{}, &entities.User{})

	DB = db

	RDB = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
