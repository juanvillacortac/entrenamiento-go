package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/juanvillacortac/entrenamiento-go/pkg/db"
	"github.com/juanvillacortac/entrenamiento-go/pkg/fetchers"
	"github.com/juanvillacortac/entrenamiento-go/pkg/queries"
)

func main() {
	godotenv.Load()

	r := gin.Default()

	db.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		params := fetchers.Params{}
		if err := c.BindQuery(&params); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		res := queries.QuerySongs(params)
		// if err != nil {
		// 	c.AbortWithError(http.StatusInternalServerError, err)
		// }
		c.JSON(http.StatusOK, res)
	})
	r.Run()
}
