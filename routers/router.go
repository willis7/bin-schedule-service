package routers

import (
    "github.com/gorilla/mux"
	"github.com/willis7/bin-schedule-service/handlers"
)

// InitRoutes initializes all the applications routes
func InitRoutes(a *handlers.App) *mux.Router {
    router := mux.NewRouter().StrictSlash(false)
    // Routes for the Schedule entity
    router = SetScheduleRoutes(router, a)
	return router
}
