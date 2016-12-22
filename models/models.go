package models

// Schedule is the canonical structure of the Bin Schedule Service
type Schedule struct {
	Postcode      string
	RecyclingDate string
	RubbishDate   string
}

// EastDevonResponse is the structured format of the response from eastdevon.gov.uk
type EastDevonResponse struct {
	Label   string `json:"label"`
	UPRN    string `json:"UPRN"`
	Result string `json:"Results"`
}
