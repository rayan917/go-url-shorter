package main

import (
	"net/http"

	"fmt"
	"hash/crc32"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

var client *redis.Client

func main() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	r := mux.NewRouter()
	r.HandleFunc("/{shorturl}", redirect).Methods("GET")
	r.HandleFunc("/", showAll).Methods("GET")
	r.HandleFunc("/", shorten).Methods("POST")
	http.ListenAndServe(":8080", r)
}

func showAll(w http.ResponseWriter, r *http.Request) {
	keys, err := client.Keys("*").Result()
	if err != nil {
		http.Error(w, "Error getting keys", http.StatusInternalServerError)
		return
	}
	fmt.Println(keys)

	for _, key := range keys {
		val, err := client.Get(key).Result()
		if err != nil {
			http.Error(w, "Error getting value for key "+key, http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%s: %s\n", key, val)
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	shorturl := mux.Vars(r)["shorturl"]
	url, err := client.Get(shorturl).Result()
	if err == redis.Nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error getting URL", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func shorten(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "Missing URL", http.StatusBadRequest)
		return
	}
	shorturl := generateShortURL(url)
	err := client.Set(shorturl, url, 0).Err()
	if err != nil {
		http.Error(w, "Error saving URL", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("http://localhost:8080/" + shorturl))
}

func generateShortURL(url string) string {
	// This is a very basic way of generating a short URL. In a real application, you would
	// want to use a more robust method of generating unique short URLs.
	return fmt.Sprintf("%x", crc32.ChecksumIEEE([]byte(url)))
}
