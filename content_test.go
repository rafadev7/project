package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

//
// Test Main
//
func XTestGETContent(t *testing.T) {

	fmt.Println("\n--- Test Content ---\n")

	w := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/api/contents", nil)
	if err != nil {
		log.Panic(err)
	}

	route.ServeHTTP(w, req)

	fmt.Println("RETURN:")
	fmt.Printf("%s", w.Body.String())

	fmt.Println("\n--- End Test Content ---\n")

}

func TestPOSTContent(t *testing.T) {

	fmt.Println("\n--- Test Content ---\n")

	c := ContentPOST{
		URL: "www.google.com",
	}
	c.Title = "Fucking shit"
	c.Description = "Its a fucking shit content"

	b, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/api/categories/internet/contents", bytes.NewBuffer(b))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	route.ServeHTTP(w, req)

	fmt.Println("RETURN:")
	fmt.Printf("%s", w.Body.String())

	fmt.Println("\n--- End Test Content ---\n")

}
