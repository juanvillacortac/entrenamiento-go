package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juanvillacortac/entrenamiento-go/pkg/api"
	"github.com/juanvillacortac/entrenamiento-go/pkg/query"
)

// GetSongs      godoc
// @Summary     Get songs array
// @Description Responds with the list of all books as JSON.
// @Tags        songs
// @Accept      json
// @Produce     json
// @Param       name   query   string  false "Song name"
// @Param       album  query   string  false "Album name"
// @Param       artist query   string  false "Artist name"
// @Param       force  query   boolean false "Bypass cache"
// @Success     200    {array} entities.Song
// @Router      /search [get]
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
