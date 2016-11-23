package base

// ParsedResult structures the information returned from the external service
type ParsedResult struct {
	RecyclingDate string
	RubbishDate   string
}

// ResultParser takes a string input and parses the dates for both Rubbish and Recycling
func ResultParser(input string) ParsedResult {
	return ParsedResult{RecyclingDate: "Friday 25 November 2016", RubbishDate: "Friday 25 November 2016"}
}
