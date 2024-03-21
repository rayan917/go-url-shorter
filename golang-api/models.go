package main

type Link struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Result struct {
	TotalLinks int    `json:"total_links"`
	Links      []Link `json:"links"`
}

type ShortURL struct {
	URL string `json:"url"`
}

type Stats struct {
	Count string `json:"count"`
}

type URLData struct {
	URL string `json:"url"`
}
