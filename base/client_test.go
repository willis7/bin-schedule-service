package base

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
)

func TestServerReturnsNon200(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	// Expect an error that matches the implementation
	payload, err := GetBinDates(ts.URL)
	expectedErr := fmt.Errorf("HTTP Response Error %d", http.StatusNotFound)
	if err.Error() != expectedErr.Error() {
		t.Errorf("GetBinDates failed: got %v expected %v for error", err, expectedErr)
	}

	if payload != nil {
		t.Errorf("GetBinDates failed: got %v expected %v for payload", payload, nil)
	}
}

//func TestServerReturnsMalformed

//func TestServerReturns200(t *testing.T) {
//	body := string(`[{"label":"label","UPRN":"uprn","Results":"result"}]`)
//	bodyBytes := []byte(body)
//	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(http.StatusOK)
//		w.Write(bodyBytes)
//	}))
//	defer ts.Close()
//
//	// Expect an error that matches the implementation
//	payload, err := GetBinDates(ts.URL)
//	expectedErr := fmt.Errorf("HTTP Response Error %d", http.StatusNotFound)
//	if err.Error() != expectedErr.Error() {
//		t.Errorf("GetBinDates failed: got %v expected %v for error", err, expectedErr)
//	}
//
//	var expectedPayload = []EastDevonResponse{EastDevonResponse{Label:"label", UPRN: "uprn", Results:"result"}}
//	if payload != nil {
//		t.Errorf("GetBinDates failed: got %v expected %v for payload", payload, expectedPayload)
//	}
//}
