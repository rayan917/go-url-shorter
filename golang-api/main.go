package main

import (
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

var client *redis.Client

func main() {
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	r := mux.NewRouter()
	r.HandleFunc("/{shorturl}", redirect).Methods("GET")
	r.HandleFunc("/", showAll).Methods("GET")
	r.HandleFunc("/", shorten).Methods("POST")
	r.HandleFunc("/{shorturl}/stats", stats).Methods("GET")
	http.ListenAndServe(":8080", r)
}
