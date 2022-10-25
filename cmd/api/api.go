package main

import (
	"github.com/gin-gonic/gin"

	"api/pkg/http"
	"api/pkg/lib/sql"
	"api/pkg/service"
)

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Status": "Ok",
		})
	})

	db := sql.NewDBConnection("./pkg/service/internal/repository/db/chinook.db")
	services := service.GetAll(db.DB)
	handlers := http.NewHandlers(services)

	r.GET("/album", handlers.HandleAlbum)

	r.Run()
}
