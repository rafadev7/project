package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

//
// Test Main
//
func XTestUser(t *testing.T) {

	fmt.Println("\n--- Test Main ---\n")

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/user/login", nil)
	if err != nil {
		log.Panic(err)
	}

	route.ServeHTTP(w, req)

	fmt.Printf("Return:\n%s\n", w.Body.String())

	fmt.Println("\n--- End Test Main ---\n")

}
