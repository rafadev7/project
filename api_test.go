package main

import (
	"fmt"
	"github.com/resourcerest/api"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

//
// Test Main
//
func XTestApi(t *testing.T) {

	fmt.Println("\n--- Test API ---\n")

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api", nil)
	if err != nil {
		log.Panic(err)
	}

	route.ServeHTTP(w, req)

	fmt.Println("RETURN:")
	fmt.Printf("%s", w.Body.String())

	fmt.Println("\n--- End Test Main ---\n")
}

//
// Test Main
//
func XTestPrint(t *testing.T) {

	fmt.Println("\n--- Test Print ---\n")

	api.PrintRoute(route)

	fmt.Println("\n--- End Test Print ---\n")

}
