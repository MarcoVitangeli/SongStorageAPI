package main

import (
	"github.com/MarcoVitangeli/SongStorageAPI/api"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	server := gin.Default()
	server.GET("/songs/:id", api.GetById)
	server.GET("/songs", api.GetAllSongs)
	server.POST("/songs/insert", api.InsertSong)

	if err := server.Run(); err != nil {
		panic(err)
	}
}
