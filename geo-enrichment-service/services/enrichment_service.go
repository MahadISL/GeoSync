package services

import (
	"log"
	"sync"

	"github.com/MahadISL/GeoSync/geo-enrichment-service/clients"
	"github.com/MahadISL/GeoSync/geo-enrichment-service/models"
)

// EnrichLocation concurrently fetches data and aggregates it.
func EnrichLocation(lat, lon float64) models.EnrichmentResponse {

	var wg sync.WaitGroup

	// Using channels to receive the results from our concurrent API calls.
	weatherChan := make(chan *clients.OWMResponse, 1)

	// We have one concurrent operation to wait for.
	wg.Add(1)

	// Launching a goroutine to fetch weather data.
	go func() {
		defer wg.Done()
		weatherData, err := clients.GetWeatherData(lat, lon)
		if err != nil {
			log.Printf("Error getting weather data: %v", err)
			weatherChan <- nil // Send nil on error
			return
		}
		weatherChan <- weatherData // Send the result to the channel.
	}()

	// Wait for all the goroutines we launched to finish.
	wg.Wait()

	// Close the channels
	close(weatherChan)

	weatherResult := <-weatherChan

	// Build the final response.
	response := models.EnrichmentResponse{}

	// Safely check if the weather result is not nil before accessing it.
	if weatherResult != nil && len(weatherResult.Weather) > 0 {
		response.Weather.Temperature = weatherResult.Main.Temp
		response.Weather.Condition = weatherResult.Weather[0].Main
	}

	return response
}
