package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/willis7/bin-schedule-service/services"
)

// App defines the application container
type App struct {
	Schedule services.Client
}

// GetSchedule returns a list of schedules for a given postcode
func (a *App)GetSchedule(w http.ResponseWriter,r *http.Request) {
	postcode := r.FormValue("postcode")
	if postcode == "" {
		http.Error(w, "MISSING_ARG_POSTCODE", http.StatusBadRequest)
		return
	}

	schedule,err := a.Schedule.Get(postcode)
	if err != nil {
		http.Error(w, "INTERNAL_ERROR", http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(schedule)
	if err != nil {
		http.Error(w, "INTERNAL_ERROR", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
