package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AvinashBarade/url-shortener/handlers"
)

func TestShortenHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/shorten?url=https://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.ShortenHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Shortened URL"
	if !contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func contains(response, expected string) bool {
	return len(response) >= len(expected) && response[:len(expected)] == expected
}
