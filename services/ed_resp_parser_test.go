package services

import (
	"bytes"
	"testing"
)

func TestResultParser(t *testing.T) {
	// When I have a result string
	result := `[{ "label": "test", "UPRN": "001", "Results": "<h2>Your recycling and food waste</h2><p class=\"boxtext-important text-centre\">Your next recycling and food waste collection will be on <em>Friday<br/>23 December 2016</em></p><p>We collect your recycling and food waste every week and your usual collection day is Friday. Your last collection was on Friday 16 December 2016.</p><h2>Your rubbish</h2><p class=\"boxtext-important text-centre\">Your next rubbish collection will be on <em>Friday<br/>23 December 2016</em></p><p>We collect your rubbish every two weeks and your usual collection day is Friday. Your last collection was on Friday 9 December 2016.</p><h2>Collections calendar</h2><p><a href=\"http://eastdevon.gov.uk/recycling-and-rubbish/when-is-my-bin-collected/future-collections-calendar/?UPRN=010000268236\" rel=\"nofollow\">View the scheduled collections for December 2016 - April 2017</a>.</p>"}]`

	// Then I expect the result to be parsed into a struct
	expected := struct {
		Postcode      string
		RecyclingDate string
		RubbishDate   string
	}{
		"fake",
		"Friday 23 December 2016",
		"Friday 23 December 2016",
	}

	// And the actual match the expected struct
	if actual := ResultParser(bytes.NewBufferString(result)); actual != expected {
		t.Errorf("ResultParser failed: expected %v, got %v", expected, actual)
	}

}
