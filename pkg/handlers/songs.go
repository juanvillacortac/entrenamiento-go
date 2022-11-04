package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juanvillacortac/entrenamiento-go/pkg/fetchers"
	"github.com/juanvillacortac/entrenamiento-go/pkg/queries"
)

func HandleSongs(c *gin.Context) {
	params := fetchers.Params{}
	if err := c.BindQuery(&params); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	songs, err := queries.QuerySongs(params, true)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, songs)
}