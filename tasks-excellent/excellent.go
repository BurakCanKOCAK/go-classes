package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Our resource struct
type SomeResource struct {
	SomeField string `json:"somefield"`
	/*
		bare in mind that field that starts with
		lowercase character is by default not exported
		and thus invisible to the JSON package
	*/
	anotherField string `json:"anotherfield"`
}

var resources []SomeResource

// Get all resources
func getAllResources(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resources)
}

// Main function
func main() {
	r := mux.NewRouter()

	resources = append(resources, SomeResource{"lorem", "foo"})
	resources = append(resources, SomeResource{"ipsum", "bar"})
	resources = append(resources, SomeResource{"dolor", "baz"})

	// registering endpoints
	r.HandleFunc("/goapi/someresources", getAllResources).Methods("GET")
	//r.HandleFunc("/goapi/someresources/{id}", getResourceByID).Methods("GET")

	// Starting the server at port 8000 with logging
	log.Fatal(http.ListenAndServe(":8000", r))
}
