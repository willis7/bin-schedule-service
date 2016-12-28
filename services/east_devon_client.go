package services

import (
	"net/http"
	"encoding/json"
	"fmt"
	"bytes"
	"net/url"

	"github.com/willis7/bin-schedule-service/models"
)

// EastDevonClient provides an implementation of the Client interface
type EastDevonClient struct {}

// Get returns an array of Schedule structs which have been returned from
func (c EastDevonClient) Get(postcode string) ([]models.Schedule, error) {
// TODO: seperate the request and the format into their own method
	// Construct the URL with params
	u := "http://eastdevon.gov.uk/addressfinder"
	vals := make(url.Values)
	vals.Set("qtype", "bins")
	vals.Set("term", fmt.Sprintf("%s", postcode))
	url := u + "?" + vals.Encode()

	// Call the East Devon API
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		statusNotFoundErr := fmt.Errorf("HTTP Response Error %d", http.StatusNotFound)
		return nil, statusNotFoundErr
	}
// ENDTODO

	// Decode top level JSON array into a slice of structs
	data := make([]models.EastDevonResponse, 0)
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		decodeErr := fmt.Errorf("Response Body Decode Error")
		return nil, decodeErr
	}

	schedules := CanonicalFormat(data, postcode)

	// return payload, nil
	return schedules, nil
}

// CanonicalFormat takes the response and
func CanonicalFormat(data []models.EastDevonResponse, postcode string) []models.Schedule {
	var schedules []models.Schedule
	for _, it := range data {
		schedule := ResultParser(bytes.NewBufferString(it.Result), postcode)
		schedules = append(schedules, schedule)
	}
	return schedules
}
