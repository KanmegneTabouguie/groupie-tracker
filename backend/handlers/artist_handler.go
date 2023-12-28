// handlers/handlers.go
package handlers

import (
	"backend/models" // Update with your actual package structure
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetArtistsHandler(c *gin.Context) {
	// Make a request to the API
	apiURL := "https://groupietrackers.herokuapp.com/api/artists"
	response, err := http.Get(apiURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from the API"})
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	// Unmarshal the JSON response into a slice of models.Artist
	var artists []models.Artist
	err = json.Unmarshal(body, &artists)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal JSON"})
		return
	}

	// Return the fetched artists as JSON
	c.JSON(http.StatusOK, artists)
}

func GetArtistByIDHandler(c *gin.Context) {
	// Get the artist ID from the URL parameters
	idParam := c.Param("id")
	artistID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid artist ID"})
		return
	}

	// Make a request to the API to get the artist by ID
	apiURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", artistID)
	response, err := http.Get(apiURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from the API"})
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	// Unmarshal the JSON response into a models.Artist
	var artist models.Artist
	err = json.Unmarshal(body, &artist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal JSON"})
		return
	}

	// Fetch and replace URL data with actual data for locations, concertDates, and relations
	artist.Locations = fetchDataFromURL(artist.Locations)
	artist.ConcertDates = fetchDataFromURL(artist.ConcertDates)
	artist.Relations = fetchDataFromURL(artist.Relations)

	// Return the fetched artist as JSON
	c.JSON(http.StatusOK, artist)
}

func fetchDataFromURL(url string) string {
	// Make a request to the specified URL
	response, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("Failed to fetch data from URL: %s", url)
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Sprintf("Failed to read response body from URL: %s", url)
	}

	// Return the data as a string
	return string(body)
}
