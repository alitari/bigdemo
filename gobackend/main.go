package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID string `json:"id,omitempty"`
}

var people []Person

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	p := Person{ID: "1"}
	people = append(people, p)
	json.NewEncoder(w).Encode(people)
}
