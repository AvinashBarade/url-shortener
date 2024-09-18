package utils

import (
	"testing"
)

// TestShortenURL checks if the ShortenURL function returns the expected length and consistent hash.
func TestShortenURL(t *testing.T) {
	originalURL := "https://example.com"
	shortenedURL := ShortenURL(originalURL)

	// Check if the length of the shortened URL is 6 characters
	if len(shortenedURL) != 6 {
		t.Errorf("expected shortened URL length of 6, got %d", len(shortenedURL))
	}

	// Check if the function is consistent (i.e., same input produces the same output)
	expectedShortened := ShortenURL(originalURL)
	if shortenedURL != expectedShortened {
		t.Errorf("expected shortened URL %s, got %s", expectedShortened, shortenedURL)
	}

	// Check different URLs produce different hashes
	anotherURL := "https://anotherexample.com"
	anotherShortened := ShortenURL(anotherURL)
	if shortenedURL == anotherShortened {
		t.Errorf("expected different shortened URLs, but got the same: %s", shortenedURL)
	}
}
