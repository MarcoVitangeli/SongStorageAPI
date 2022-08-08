package api

import (
	"database/sql"
	"github.com/MarcoVitangeli/SongStorageAPI/db"
	"github.com/MarcoVitangeli/SongStorageAPI/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetById(c *gin.Context) {
	conn := db.CreateConnexion()
	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	id := c.Param("id")

	row := conn.QueryRow("SELECT name, author, releaseDate FROM song WHERE id = ?", id)

	song := model.Song{}
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id must be a number",
		})
		return
	}
	song.Id = int(intId)
	err = row.Scan(&song.Name, &song.Author, &song.ReleaseDate)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "id does not match an existing song",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, song)
}

func GetAllSongs(c *gin.Context) {
	conn := db.CreateConnexion()
	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	rows, err := conn.Query("SELECT id, name, author, releaseDate FROM song")
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	if err != nil {
		panic(err)
	}

	var songs []model.Song

	for rows.Next() {
		var song model.Song
		err = rows.Scan(&song.Id, &song.Name, &song.Author, &song.ReleaseDate)

		if err != nil {
			panic(err)
		}
		songs = append(songs, song)
	}

	c.IndentedJSON(http.StatusOK, songs)
}

func InsertSong(c *gin.Context) {
	conn := db.CreateConnexion()
	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	var song model.Song
	if err := c.BindJSON(&song); err != nil {
		panic(err)
	}

	res, err := conn.Exec("INSERT INTO song(name, author, releaseDate) VALUES(?, ?, ?)", song.Name, song.Author, song.ReleaseDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error in the song parameters",
		})
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
