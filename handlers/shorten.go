package handlers

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/AvinashBarade/url-shortener/utils"
)

var urlStore = make(map[string]string) // In-memory store for URLs
var mu sync.Mutex                      // Mutex for concurrency control

// ShortenHandler handles the URL shortening requests
func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	originalURL := r.URL.Query().Get("url")
	if originalURL == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	mu.Lock()
	shortURL, exists := urlStore[originalURL]
	if !exists {
		shortURL = utils.ShortenURL(originalURL)
		urlStore[originalURL] = shortURL
	}
	mu.Unlock()

	fmt.Fprintf(w, "Shortened URL: %s", shortURL)
}
