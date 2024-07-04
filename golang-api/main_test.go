package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestShowAll(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	client = redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	s.Set("shorturl1", "http://example1.com")
	s.Set("shorturl2", "http://example2.com")
	s.Set("count:shorturl1", "5")

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/", showAll)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var result Result
	json.Unmarshal(rr.Body.Bytes(), &result)

	assert.Equal(t, 2, result.TotalLinks)

	expectedLinks := []Link{
		{Key: "shorturl1", Value: "http://example1.com"},
		{Key: "shorturl2", Value: "http://example2.com"},
	}
	assert.ElementsMatch(t, expectedLinks, result.Links)
}

func TestRedirect(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	client = redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	s.Set("shorturl", "http://example.com")

	req, _ := http.NewRequest("GET", "/shorturl", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/{shorturl}", redirect)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusFound, rr.Code)
	assert.Equal(t, "http://example.com", rr.Header().Get("Location"))
}

func TestShorten(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	client = redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	data := URLData{URL: "http://example.com"}
	jsonBody, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonBody))
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/", shorten)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var result ShortURL
	json.Unmarshal(rr.Body.Bytes(), &result)
}

func TestStats(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	client = redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	s.Set("count:shorturl", "10")

	req, _ := http.NewRequest("GET", "/shorturl/stats", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/{shorturl}/stats", stats)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var result Stats
	json.Unmarshal(rr.Body.Bytes(), &result)
	assert.Equal(t, "10", result.Count)
}
