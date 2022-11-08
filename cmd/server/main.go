package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/juanvillacortac/entrenamiento-go/docs"
	"github.com/juanvillacortac/entrenamiento-go/pkg/db"
	"github.com/juanvillacortac/entrenamiento-go/pkg/handlers"
	"github.com/juanvillacortac/entrenamiento-go/pkg/middlewares"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title    Songs Indexer
// @version  1.0
// @BasePath /api/v1

func main() {
	godotenv.Load()

	r := gin.Default()

	db.ConnectDatabase()

	authMiddleware := middlewares.AuthMiddleware(r)

	docs.SwaggerInfo.BasePath = "/api/v1"
	api := r.Group("/api/v1")
	api.Use(authMiddleware)
	{
		api.GET("/search", handlers.HandleSongs)
	}
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}
