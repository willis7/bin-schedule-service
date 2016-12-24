package main

import (
	"log"
	"net/http"

	"github.com/willis7/bin-schedule-service/handlers"
	"github.com/willis7/bin-schedule-service/services"
	"github.com/willis7/bin-schedule-service/routers"
	"github.com/urfave/negroni"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	app := &handlers.App{Schedule: services.EastDevonClient{}}

	// Get the mux router object
    router := routers.InitRoutes(app)

	// Create a negroni instance
    n := negroni.Classic()
    n.UseHandler(router)

	server := &http.Server{
        Addr:    ":" + port,
        Handler: n,
    }
    log.Println("Listening...")
	log.Fatal(server.ListenAndServe())
}
