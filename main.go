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
	{ID: "4", Title: "Slipknot", Artist: "Slipknot", Year: 1999},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	r := gin.Default()

	r.GET("/albums", func(ctx *gin.Context) {
		getAlbums(ctx)
	})

	r.Run()
}
