package handlers

import (
	"net/http"
	"sync"
)

var muRedirect sync.Mutex // Mutex for redirection

// RedirectHandler handles the redirection from shortened URL to original URL
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]

	muRedirect.Lock()
	originalURL := findOriginalURL(shortURL)
	muRedirect.Unlock()

	if originalURL != "" {
		http.Redirect(w, r, originalURL, http.StatusFound)
	} else {
		http.NotFound(w, r)
	}
}

// Helper function to find the original URL for a given short URL
func findOriginalURL(shortURL string) string {
	for originalURL, short := range URLStore {
		if short == shortURL {
			return originalURL
		}
	}
	return ""
}
