package main

import (
	"fmt"
	"hash/crc32"
	"testing"
)

func TestGenerateShortURL(t *testing.T) {
	url := "https://www.example.com/"
	expected := fmt.Sprintf("%x", crc32.ChecksumIEEE([]byte(url)))

	result := generateShortURL(url)

	if result != expected {
		t.Errorf("generateShortURL(%s) = %s; want %s", url, result, expected)
	}
}
