package routers

import (
	"github.com/gorilla/mux"
	"github.com/willis7/bin-schedule-service/handlers"
)

// SetScheduleRoutes defines a set of routes for the schedules resources
func SetScheduleRoutes(router *mux.Router, a *handlers.App) *mux.Router {
	router.HandleFunc("/schedules", a.GetSchedule).Methods("GET")

	return router
}
