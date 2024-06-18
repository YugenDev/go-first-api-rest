package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

var albums = []Album{
	{ID: "1", Title: "Laterlus", Artist: "Tool", Year: 2001},
	{ID: "2", Title: "Desire, I Want To Turn Into You", Artist: "Caroline Polachek", Year: 2023},
	{ID: "3", Title: "Angel Dust", Artist: "Faith No More", Year: 1992},
	{ID: "4", Title: "MIguel cacorro", Artist: "Slipknot", Year: 1999},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, albums)
}

func getAlbumByID(c *gin.Context){
	id := c.Param("id")

	for _, album := range albums{
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
		
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album no encontrado"})
}

func main() {
	r := gin.Default()

	r.GET("/albums", getAlbums)
	r.POST("/albums", postAlbums)
	r.GET("/albums/:id", getAlbumByID)

	r.Run()
}
