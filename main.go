package main

import (
	"net/http"
	"strings"

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
	r.HandleFunc("/{shorturl}/stats", stats).Methods("GET")
	http.ListenAndServe(":8080", r)
}

func showAll(w http.ResponseWriter, r *http.Request) {
	keys, err := client.Keys("*").Result()
	if err != nil {
		http.Error(w, "Error getting keys", http.StatusInternalServerError)
		return
	}

	count := 0

	for _, key := range keys {

		if !strings.Contains(key, "count") {
			count++
			fmt.Fprintf(w, "Total links: %d\n", count)
			val, err := client.Get(key).Result()
			if err != nil {
				http.Error(w, "Error getting value for key "+key, http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(w, "%s: %s\n", key, val)
		}
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
	client.Incr("count:" + shorturl)
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
	client.Set("count:"+shorturl, 0, 0)
	w.Write([]byte("http://localhost:8080/" + shorturl))
}

func stats(w http.ResponseWriter, r *http.Request) {
	shorturl := mux.Vars(r)["shorturl"]
	count, err := client.Get("count:" + shorturl).Result()

	if err != nil {
		http.Error(w, "Error getting count", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Count: " + count))
}

func generateShortURL(url string) string {
	return fmt.Sprintf("%x", crc32.ChecksumIEEE([]byte(url)))
}
