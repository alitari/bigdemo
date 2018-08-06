package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

// Message here comes a comment
type Message struct {
	ID   string `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
}

var messages []Message

var redisClient *redis.Client

// our main function
func main() {

	redisHost := os.Getenv("REDIS_HOST")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, "6379"),
		Password: redisPassword,
		DB:       0, // use default DB
	})

	router := mux.NewRouter()
	router.HandleFunc("/messages/{id}", GetMessage).Methods("GET")

	router.HandleFunc("/messages", GetMessages).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

// GetMessage for ID
func GetMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	message := message(params["id"])
	if message.ID == "error" {
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(message)
}

func message(id string) Message {
	te, err := redisClient.Get(id).Result()
	var message Message
	if err != nil {
		message = Message{ID: "error", Text: err.Error()}
	} else {
		message = Message{ID: id, Text: te}
	}
	return message
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	words, hasQueryParam := r.URL.Query()["word"]
	var messages = []Message{}
	if !hasQueryParam {
		w.WriteHeader(http.StatusBadRequest)
		messages = []Message{Message{ID: "error", Text: "need query param 'word'"}}
	} else {
		for _, wrd := range words {
			ids, err := redisClient.LRange(wrd, 0, 10).Result()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				messages = []Message{Message{ID: "error", Text: err.Error()}}
			} else {
				for _, id := range ids {
					messages = append(messages, message(id))
				}
			}
		}
	}
	json.NewEncoder(w).Encode(messages)
}
