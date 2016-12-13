package main

import (
	"fmt"
	. "github.com/willis7/bin-schedule-service/base"
	"net/url"
)

var postcode = "EX5 3EL"

func main() {
	// Construct the URL with params
	u := "http://eastdevon.gov.uk/addressfinder"
	vals := make(url.Values)
	vals.Set("qtype", "bins")
	vals.Set("term", fmt.Sprintf("%s", postcode))
	fullUrl := u + "?" + vals.Encode()

	response, err := GetBinDates(fullUrl)
	if err != nil {
		fmt.Sprintf("GetBinDates: %v", err)
	}
	fmt.Printf("%#v", response)
}
