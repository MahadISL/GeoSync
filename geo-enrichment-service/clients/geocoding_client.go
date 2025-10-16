package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// NominatimResponse defines the structure for the relevant parts of the Nominatim API response.
type NominatimResponse struct {
	DisplayName string `json:"display_name"`
	Address     struct {
		City    string `json:"city"`
		Country string `json:"country"`
	} `json:"address"`
}

// GetLocationName performs a reverse geocode lookup to find a location name.
func GetLocationName(lat, lon float64) (*NominatimResponse, error) {
	// Construct the API URL.
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f", lat, lon)

	// Create a new HTTP client.
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request for Nominatim: %w", err)
	}

	req.Header.Set("User-Agent", "GeoSync-API")

	// Execute the request.
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request to Nominatim: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Nominatim API returned non-200 status: %s", resp.Status)
	}

	// Decode the JSON response.
	var nominatimResponse NominatimResponse
	if err := json.NewDecoder(resp.Body).Decode(&nominatimResponse); err != nil {
		return nil, fmt.Errorf("error decoding Nominatim response: %w", err)
	}

	return &nominatimResponse, nil
}
