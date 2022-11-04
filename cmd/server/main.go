package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/juanvillacortac/entrenamiento-go/pkg/db"
	"github.com/juanvillacortac/entrenamiento-go/pkg/handlers"
)

func main() {
	godotenv.Load()

	r := gin.Default()

	db.ConnectDatabase()

	r.GET("/", handlers.HandleSongs)
	r.Run()
}
