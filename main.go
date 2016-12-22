package main

import (
	"log"
	"net/http"

	"github.com/willis7/bin-schedule-service/handlers"
	"github.com/willis7/bin-schedule-service/services"
	"github.com/willis7/bin-schedule-service/routers"
	"github.com/urfave/negroni"
)

func main() {

	app := &handlers.App{Schedule: services.EastDevonClient{}}

	// Get the mux router object
    router := routers.InitRoutes(app)

	// Create a negroni instance
    n := negroni.Classic()
    n.UseHandler(router)

	server := &http.Server{
        Addr:    ":8080",
        Handler: n,
    }
    log.Println("Listening...")
	log.Fatal(server.ListenAndServe())
}
