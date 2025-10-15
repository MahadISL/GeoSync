package main

import (
	"github.com/MahadISL/GeoSync/geo-enrichment-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/health", handlers.HealthCheckHandler)
	router.POST("/enrich", handlers.EnrichHandler)

	router.Run(":8081")
}
