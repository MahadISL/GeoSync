package handlers

import (
	"github.com/MahadISL/GeoSync/geo-enrichment-service/models"
	"github.com/MahadISL/GeoSync/geo-enrichment-service/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func EnrichHandler(c *gin.Context) {

	var request models.EnrichmentRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	finalResponse := services.EnrichLocation(request.Latitude, request.Longitude)

	c.JSON(http.StatusOK, finalResponse)
}

func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}
