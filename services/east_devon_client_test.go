package services

import (
	"testing"
	"gopkg.in/jarcoal/httpmock.v1"
	"github.com/willis7/bin-schedule-service/models"
	"net/http"
	"reflect"
)

var c = EastDevonClient{}
var fixture = `[{ "label": "test", "UPRN": "001", "Results": "<h2>Your recycling and food waste</h2><p class=\"boxtext-important text-centre\">Your next recycling and food waste collection will be on <em>Friday<br/>23 December 2016</em></p><p>We collect your recycling and food waste every week and your usual collection day is Friday. Your last collection was on Friday 16 December 2016.</p><h2>Your rubbish</h2><p class=\"boxtext-important text-centre\">Your next rubbish collection will be on <em>Friday<br/>23 December 2016</em></p><p>We collect your rubbish every two weeks and your usual collection day is Friday. Your last collection was on Friday 9 December 2016.</p><h2>Collections calendar</h2><p><a href=\"http://eastdevon.gov.uk/recycling-and-rubbish/when-is-my-bin-collected/future-collections-calendar/?UPRN=010000268236\" rel=\"nofollow\">View the scheduled collections for December 2016 - April 2017</a>.</p>"}]`

func TestEastDevonClient(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		description string
		responder httpmock.Responder
		expectedSchedules []models.Schedule
		expectedError error
	} {
		// #1
		{
			description:   "east devon api success",
			responder:     httpmock.NewStringResponder(http.StatusOK, fixture),
			expectedSchedules: []models.Schedule{models.Schedule{Postcode: "fake", RecyclingDate: "Friday 23 December 2016", RubbishDate: "Friday 23 December 2016"}},
			expectedError: nil,
		},
		// TODO: add test for a different postcode
	}

	for _, tc := range tests {
		httpmock.RegisterResponder("GET", "http://eastdevon.gov.uk/addressfinder?qtype=bins&term=fake", tc.responder)

		r, err := c.Get("fake")

		if !reflect.DeepEqual(r, tc.expectedSchedules) {
			t.Errorf("client returned wrong schedules: got %v want %v",
				r, tc.expectedSchedules)
		}

		if err != tc.expectedError {
			t.Errorf("client returned wrong schedules: got %v want %v",
				err, tc.expectedError)
		}

		httpmock.Reset()
	}
}
