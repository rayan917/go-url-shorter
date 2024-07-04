package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

func showAll(w http.ResponseWriter, r *http.Request) {
	keys, err := client.Keys("*").Result()
	if err != nil {
		http.Error(w, "Error getting keys", http.StatusInternalServerError)
		return
	}

	var links []Link

	for _, key := range keys {
		if !strings.Contains(key, "count") {
			val, err := client.Get(key).Result()
			if err != nil {
				http.Error(w, "Error getting value for key "+key, http.StatusInternalServerError)
				return
			}
			links = append(links, Link{Key: key, Value: val})
		}
	}

	result := Result{
		TotalLinks: len(links),
		Links:      links,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
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
	http.Redirect(w, r, url, http.StatusFound)
}

func shorten(w http.ResponseWriter, r *http.Request) {
	var data URLData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Error parsing JSON body", http.StatusBadRequest)
		return
	}

	if data.URL == "" {
		http.Error(w, "Missing URL", http.StatusBadRequest)
		return
	}

	shorturl := generateShortURL(data.URL)
	err = client.Set(shorturl, data.URL, 0).Err()
	if err != nil {
		http.Error(w, "Error saving URL", http.StatusInternalServerError)
		return
	}
	client.Set("count:"+shorturl, 0, 0)

	result := ShortURL{
		URL: "http://localhost:8080/" + shorturl,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func stats(w http.ResponseWriter, r *http.Request) {
	shorturl := mux.Vars(r)["shorturl"]
	count, err := client.Get("count:" + shorturl).Result()

	if err != nil {
		http.Error(w, "Error getting count", http.StatusInternalServerError)
		return
	}

	result := Stats{
		Count: count,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
