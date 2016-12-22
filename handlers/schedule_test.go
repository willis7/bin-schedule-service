package handlers

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/willis7/bin-schedule-service/models"
	"errors"
)

// Mock for the client dependencies
type TestClient struct {
	Schedule []models.Schedule
	Err      error
}

func (c TestClient) Get(string) ([]models.Schedule, error) {
	return c.Schedule, c.Err
}

func TestGetSchedule(t *testing.T) {

	tests := []struct {
		description        string
		testClient         *TestClient
		url                string
		expectedStatusCode int
		expectedBody       string
	}{
		// #1
		{
			description:        "missing argument user",
			testClient:        &TestClient{},
			url:                "/schedule",
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "MISSING_ARG_POSTCODE\n",
		},
		// #2
		{
			description: "error getting schedule",
			testClient: &TestClient{
				Schedule: []models.Schedule{},
				Err:      errors.New("fake test error"),
			},
			url:                "/repos?postcode=fakecode",
			expectedStatusCode: http.StatusInternalServerError,
			expectedBody:       "INTERNAL_ERROR\n",
		},
		// #3
		{
			description: "no schedules found",
			testClient: &TestClient{
				Schedule: []models.Schedule{},
				Err:      nil,
			},
			url:                "/repos?postcode=fakecode",
			expectedStatusCode: http.StatusOK,
			expectedBody:       `[]`,
		},
		// #4
		{
			description: "successful query",
			testClient: &TestClient{
				Schedule: []models.Schedule{
					models.Schedule{Postcode: "fakecode", RecyclingDate: "fake recycling", RubbishDate: "fake rubbish"},
				},
				Err:      nil,
			},
			url:                "/repos?postcode=fakecode",
			expectedStatusCode: http.StatusOK,
			expectedBody:       `[{"Postcode":"fakecode","RecyclingDate":"fake recycling","RubbishDate":"fake rubbish"}]`,
		},
	}

	for _, tc := range tests {
		app := &App{Schedule: tc.testClient}

		req, err := http.NewRequest("GET", tc.url, nil)
		if err != nil {
			t.Error("NewRequest: failed")
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		app.GetSchedule(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != tc.expectedStatusCode {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, tc.expectedStatusCode)
		}

		if body := rr.Body.String(); body != tc.expectedBody {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), tc.expectedBody)
		}
	}
}
