package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juanvillacortac/entrenamiento-go/pkg/fetchers"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		res, err := fetchers.FetchFromItunes(c.Query("q"))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, res.Transform())
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
