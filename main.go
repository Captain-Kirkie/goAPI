package main

// package import path is module path joined with subDirectory
import (
	"fmt"
	album "qpi/web-service-gin/models"
	"net/http"
    "github.com/gin-gonic/gin"
)

const PORT = 8080;
const HOSE = "localhost:";
var albumsArr = []album.Album{
		{ ID: 1, Title : "To Late for Edielweiss", Artist: "Tallest", Price: 54.32},
		{ ID: 2, Title : "Can't Wake Up", Artist: "Shakey Graves", Price: 100.12},
		{ ID: 3, Title : "2 Be Loved", Artist: "Lizzo", Price: 454.23},
	}


func main() {
	router := gin.Default();
	router.GET("/albums", getAlbums);
	router.Run("")

	fmt.Println("server running at %v", PORT)
}


func getAlbums (c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albumsArr)
}