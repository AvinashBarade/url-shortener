package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AvinashBarade/url-shortener/handlers"
)

func TestRedirectHandler(t *testing.T) {
	// Add a shortened URL to the in-memory store
	handlers.URLStore["https://example.com"] = "abc123"

	req, err := http.NewRequest("GET", "/abc123", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.RedirectHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusFound)
	}

	if location := rr.Header().Get("Location"); location != "https://example.com" {
		t.Errorf("handler returned wrong location header: got %v want %v", location, "https://example.com")
	}
}
