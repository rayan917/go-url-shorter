package main

import (
	"fmt"
	"hash/crc32"
)

func generateShortURL(url string) string {
	return fmt.Sprintf("%x", crc32.ChecksumIEEE([]byte(url)))
}
