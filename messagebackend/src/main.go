package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

// Message here comes a comment
type Message struct {
	ID           string `json:"id,omitempty"`
	Text         string `json:"text,omitempty"`
	Author       string `json:"author,omitempty"`
	CreationTime string `json:"creationTime,omitempty"`
}

type Instant struct {
	EpochSecond int64 `json:"epochSecond,omitempty"`
	Nano        int64 `json:"nano,omitempty"`
}

type dataAccess interface {
	getData(key string) (data []string, err error)
	getDataCount() (count int, err error)
	queryData(query string) (keys []string, err error)
}

type redisDataAccess struct {
	redisClient *redis.Client
}

func (r *redisDataAccess) getData(key string) (data []string, err error) {
	return r.redisClient.LRange(key, 0, 2).Result()
}
func (r *redisDataAccess) getDataCount() (count int, err error) {
	cd, err := r.redisClient.Get("message_id").Result()
	var cnt = 0
	if err == nil {
		cnt, _ = strconv.Atoi(cd)
	}
	return cnt, err
}
func (r *redisDataAccess) queryData(query string) (keys []string, err error) {
	return r.redisClient.LRange(query, 0, 10).Result()
}

var messages []Message

var dataAcc dataAccess

// our main function
func main() {
	dataAcc = &redisDataAccess{redisClient: createRedisClient()}
	serv()
}

func serv() error {
	router := mux.NewRouter()
	router.HandleFunc("/messages/{id}", GetMessage).Methods("GET")

	router.HandleFunc("/messages", GetMessages).Methods("GET")

	listenPort := fmt.Sprintf(":%s", "8000")
	fmt.Printf("listening: %s\n", listenPort)
	return http.ListenAndServe(listenPort, router)
}

func createRedisClient() *redis.Client {
	redisHost := os.Getenv("REDIS_HOST")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	fmt.Printf("redis host: %s\n", redisHost)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, "6379"),
		Password: redisPassword,
		DB:       0, // use default DB
	})
	return redisClient
}

// GetMessage for ID
func GetMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	var mess Message
	if id == "count" {
		count, err := dataAcc.getDataCount()
		if err != nil {
			mess = Message{ID: "error", Text: err.Error(), Author: "no"}
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			mess = Message{ID: id, Text: strconv.Itoa(count), Author: "no"}
		}
	} else {
		mess = message(id)
		if mess.ID == "error" {
			w.WriteHeader(http.StatusInternalServerError)
		}
		if mess.ID == "not found" {
			w.WriteHeader(http.StatusNotFound)
		}
	}
	json.NewEncoder(w).Encode(mess)
}

func message(id string) Message {
	me, err := dataAcc.getData(id)
	var message Message
	if err != nil {
		message = Message{ID: "error", Text: err.Error(), Author: "no", CreationTime: strconv.FormatInt(time.Now().Unix(), 10)}
	} else {
		if len(me) > 0 {
			var instant Instant
			err := json.Unmarshal([]byte(me[2]), &instant)
			if err != nil {
				message = Message{ID: "error", Text: err.Error(), Author: "no", CreationTime: strconv.FormatInt(time.Now().Unix(), 10)}
			}
			tiStr := strconv.FormatInt(instant.EpochSecond, 10)
			message = Message{ID: id, Text: me[0], Author: me[1], CreationTime: tiStr}
		} else {
			message = Message{ID: id, Text: "not found", Author: "no", CreationTime: strconv.FormatInt(time.Now().Unix(), 10)}
		}
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
			ids, err := dataAcc.queryData(wrd)
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
