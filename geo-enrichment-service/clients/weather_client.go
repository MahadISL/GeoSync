package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// OWMResponse is the structure for the relevant parts of the OpenWeatherMap API response.
type OWMResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Main string `json:"main"`
	} `json:"weather"`
}

// GetWeatherData fetches weather data for the given coordinates.
func GetWeatherData(lat, lon float64) (*OWMResponse, error) {

	// We will set this variable when we run our application.
	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("OPENWEATHERMAP_API_KEY environment variable not set")
	}

	// Construct the API URL. Note the units=metric to get Celsius.
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=metric", lat, lon, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request to OpenWeatherMap: %w", err)
	}
	defer resp.Body.Close() // closing connection

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("OpenWeatherMap API returned non-200 status: %s", resp.Status)
	}

	// Decode the JSON response into our struct.
	var owmResponse OWMResponse
	if err := json.NewDecoder(resp.Body).Decode(&owmResponse); err != nil {
		return nil, fmt.Errorf("error decoding OpenWeatherMap response: %w", err)
	}

	return &owmResponse, nil
}
