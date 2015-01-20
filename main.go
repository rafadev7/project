package main

import (
	"fmt"
	"github.com/resourcerest/api"
	"log"
)

var Env struct {
	Url        string
	Port       int
	Production bool
}

var route *api.Route

func init() {
	if false { // martini.Env == "production"
		log.Println("Server in production")
		Env.Port = 8000
		Env.Url = "http://tur.ma/"
		Env.Production = true
	} else {
		log.Println("Server in development")
		Env.Port = 8000
		Env.Url = fmt.Sprintf("http://localhost:%d/", Env.Port)
		Env.Production = true
	}

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
	api.PrintResource(resource)
	api.PrintRoute(route)
}

func main() {
	//server.Run(":8080")
}
