package models

type EnrichmentRequest struct {
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}

// EnrichmentResponse defines the structure of the JSON we will send back.
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
