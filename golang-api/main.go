package main

import (
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/handlers"
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

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r))

}
