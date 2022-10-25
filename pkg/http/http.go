package http

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"api/pkg/service"
)

type Handlers interface {
	HandleAlbum(c *gin.Context)
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
		if err != nil {
			panic(err)
		}

		album, err := h.services.Album.Get(i)
		if err != nil {
			panic(err)
		}

		c.JSON(200, album)
	}
}