package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type EnrichmentRequest struct {
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}
type EnrichmentResponse struct {
	Weather WeatherData `json:"weather"`
	Places  []PlaceData `json:"places"`
}
type WeatherData struct {
	Temperature float64 `json:"temperature"`
	Condition   string  `json:"condition"`
}
type PlaceData struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

func EnrichHandler(c *gin.Context) {

	var request EnrichmentRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Mocked response data.
	mockedResponse := EnrichmentResponse{
		Weather: WeatherData{
			Temperature: 15.5,
			Condition:   "Sunny",
		},
		Places: []PlaceData{
			{Name: "Eiffel Tower", Category: "Tourist Attraction"},
			{Name: "Louvre Museum", Category: "Museum"},
		},
	}

	c.JSON(http.StatusOK, mockedResponse)
}

func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}
