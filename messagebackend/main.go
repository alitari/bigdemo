package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Message here comes a comment
type Message struct {
	ID   string `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
}

var messages []Message

// our main function
func main() {
	messages = append(messages, Message{ID: "1", Text: " first message"})
	messages = append(messages, Message{ID: "2", Text: " second message"})
	router := mux.NewRouter()
	router.HandleFunc("/messages", GetMessages).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// GetMessages here comes a comment
func GetMessages(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(messages)
}
