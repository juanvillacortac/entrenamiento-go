package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/juanvillacortac/entrenamiento-go/pkg/db"
	"github.com/juanvillacortac/entrenamiento-go/pkg/handlers"
	"github.com/juanvillacortac/entrenamiento-go/pkg/middlewares"
)

func main() {
	godotenv.Load()

	r := gin.Default()

	db.ConnectDatabase()

	authMiddleware := middlewares.AuthMiddleware(r)

	api := r.Group("/api")
	api.Use(authMiddleware)
	{
		api.GET("/search", handlers.HandleSongs)
	}
	r.Run()
}
