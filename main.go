package main

import (
	"flag"
	"fmt"
	"github.com/resourcerest/api"
	"log"
	"net/http"
)

var route *api.Route

// Development Env
var env = devEnv
var prod bool

func init() {

	flag.BoolVar(&prod, "prod", false, "Production? True or False.")

	resource, err := api.NewResource(Api{
		Version: 0,
	})
	if err != nil {
		log.Fatalf("Server Fatal: %s\n", err)
	}

	route, err = api.NewRoute(resource)
	if err != nil {
		log.Fatalf("Server Fatal: %s\n", err)
	}

	// Print TESTS
	//api.PrintResource(resource)
	//api.PrintRoute(route)
}

func main() {
	flag.Parse()
	if prod {
		env = prodEnv
	}

	// Starting de HTTP server
	log.Println("Starting HTTP server in " + env.Url + " ...")
	if err := http.ListenAndServe(fmt.Sprintf(":%d", env.Port), route); err != nil {
		log.Fatalf("Server Fatal: %s\n", err)
	}
}
