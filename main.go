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

func main() {
	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "fokiu",
		})
	})

	r.Run()
}
