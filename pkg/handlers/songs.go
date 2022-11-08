package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juanvillacortac/entrenamiento-go/pkg/api"
	"github.com/juanvillacortac/entrenamiento-go/pkg/query"
)

func HandleSongs(c *gin.Context) {
	params := api.Params{}
	_, force := c.GetQuery("force")
	if err := c.BindQuery(&params); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	songs, err := query.QuerySongs(params, !force)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, songs)
}
