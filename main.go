package main

// package import path is module path joined with subDirectory
import (
	"fmt"
	"net/http"
	album "qpi/web-service-gin/models"

	"github.com/gin-gonic/gin"
)
/**
	Server specs.
*/
const PORT = "8080";
const HOST = "localhost";
var serverEndpoint = fmt.Sprintf("%v:%v",HOST, PORT);

var albumsArr = []album.Album{
		{ ID: 1, Title : "To Late for Edielweiss", Artist: "Tallest", Price: 54.32},
		{ ID: 2, Title : "Can't Wake Up", Artist: "Shakey Graves", Price: 100.12},
		{ ID: 3, Title : "2 Be Loved", Artist: "Lizzo", Price: 454.23},
	}


func main() {
	router := gin.Default();
	router.GET("/albums", getAlbums);
	router.Run(serverEndpoint);
	fmt.Println("server running at %v", serverEndpoint)
}


func getAlbums (c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albumsArr)
}