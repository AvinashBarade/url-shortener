package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// ShortenURL generates a shortened URL hash based on the original URL

func ShortenURL(originalURL string) string {
	hash := md5.Sum([]byte(originalURL))   // Create an MD5 hash of the original URL
	return hex.EncodeToString(hash[:])[:6] // Use the first 6 characters as the shortened URL
}
