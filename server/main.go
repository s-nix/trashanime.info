package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/anime", getAnime)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

type Anime struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Genre string `json:"genre"`
}

var anime = []Anime{
	{ID: "1", Title: "Naruto", Genre: "Action"},
	{ID: "2", Title: "One Piece", Genre: "Adventure"},
	{ID: "3", Title: "Death Note", Genre: "Thriller"},
}

func getAnime(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, anime)
}
