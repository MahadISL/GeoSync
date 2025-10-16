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

	// Create a channel for each concurrent API call.
	weatherChan := make(chan *clients.OWMResponse, 1)
	geoChan := make(chan *clients.NominatimResponse, 1)

	// Two concurrent operation to wait for.
	wg.Add(2)

	// --- Goroutine 1: Fetch Weather Data ---
	go func() {
		log.Println("Fetching weather data...")
		defer wg.Done()
		weatherData, err := clients.GetWeatherData(lat, lon)
		if err != nil {
			log.Printf("Error getting weather data: %v", err)
			weatherChan <- nil // Send nil on error
			return
		}
		weatherChan <- weatherData // Send the result to the channel.
	}()

	// --- Goroutine 2: Fetch Location Name ---
	go func() {
		log.Println("Fetching location name...")
		defer wg.Done() // Decrement the WaitGroup counter when this goroutine finishes.
		geoData, err := clients.GetLocationName(lat, lon)
		if err != nil {
			log.Printf("Error getting location name: %v", err)
			geoChan <- nil
			return
		}
		geoChan <- geoData
	}()

	// Wait for all the goroutines we launched to finish.
	wg.Wait()

	// Close the channels.
	close(weatherChan)
	close(geoChan)

	// Read the results from the channels.
	weatherResult := <-weatherChan
	geoResult := <-geoChan

	// Build the final response.
	response := models.EnrichmentResponse{
		Places: []models.PlaceData{}, // Initialize as an empty slice to avoid 'null' in JSON.
	}

	// Safely check if the weather result is not nil before accessing it.
	if weatherResult != nil && len(weatherResult.Weather) > 0 {
		response.Weather.Temperature = weatherResult.Main.Temp
		response.Weather.Condition = weatherResult.Weather[0].Main
	}

	if geoResult != nil {
		response.Places = append(response.Places, models.PlaceData{
			Name:     geoResult.DisplayName,
			Category: "Location Address",
		})
	}

	return response
}
