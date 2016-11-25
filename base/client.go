package base

import (
	"net/http"
	"encoding/json"
	"fmt"
)

// EastDevonResponse is the structured format of the response from eastdevon.gov.uk
type EastDevonResponse struct {
	Label   string `json:"label"`
	UPRN    string `json:"UPRN"`
	Results string `json:"Results"`
}

// GetBinDates returns an array of EastDevonResponses which match the postcode
func GetBinDates(url string) ([]EastDevonResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		non200Err := fmt.Errorf("HTTP Response Error %d", http.StatusNotFound)
		return nil, non200Err
	}

	// Decode top level JSON array into a slice of structs
	payload := make([]EastDevonResponse, 0)
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		decodeErr := fmt.Errorf("Response Body Decode Error")
		return nil, decodeErr
	}
	return payload, nil
}
