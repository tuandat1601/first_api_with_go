package main

import (
	
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
func addAlbum(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusOK, newAlbum)
}
func getAlbumById( c *gin.Context){
	id:= c.Param("id")
	for _, a:= range albums{
		if a.ID ==id{
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
func updateAlbumById(c *gin.Context) {
	id := c.Param("id")
    
	// Loop through the list of albums to find the album by ID.
	for i, a := range albums {
	    if a.ID == id {
		// Create a new album to store the changes.
		var updatedAlbum album
    
		// Bind the received JSON to the new album.
		if err := c.BindJSON(&updatedAlbum); err != nil {
		    return
		}
    
		// Update only the fields that are present in the received JSON.
		if updatedAlbum.Title != "" {
		    albums[i].Title = updatedAlbum.Title
		}
		if updatedAlbum.Artist != "" {
		    albums[i].Artist = updatedAlbum.Artist
		}
		if updatedAlbum.Price != 0 {
		    albums[i].Price = updatedAlbum.Price
		}
    
		// Respond with the updated album.
		c.IndentedJSON(http.StatusOK, albums[i])
		return
	    }
	}
    
	// If no matching album is found, respond with a 404 Not Found status.
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
    }
 
    
    
    
func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.POST("/albums", addAlbum)
	router.GET("/albums/:id",getAlbumById)
	router.PATCH("/albums/:id",updateAlbumById)
	router.Run("localhost:8080")
}
