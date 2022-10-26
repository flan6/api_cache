package http

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"api/pkg/service"
)

type Handlers interface {
	HandleAlbum(c *gin.Context)
	HandleArtist(c *gin.Context)
}

type handlers struct {
	services service.All
}

func NewHandlers(services service.All) Handlers {
	return handlers{
		services: services,
	}
}

func (h handlers) HandleAlbum(c *gin.Context) {
	id := c.Query("id")
	if id != "" {
		i, err := strconv.Atoi(id)
		panicErr(err)

		album, err := h.services.Album.Get(i)
		panicErr(err)

		c.JSON(200, album)
	}
}

func (h handlers) HandleArtist(c *gin.Context) {
	id := c.Query("id")
	if id != "" {
		i, err := strconv.Atoi(id)
		panicErr(err)

		artist, err := h.services.Artist.Get(i)
		panicErr(err)

		c.JSON(200, artist)
	}
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
