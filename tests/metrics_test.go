package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AvinashBarade/url-shortener/handlers"
)

func TestMetricsHandler(t *testing.T) {
	// Add URLs to the in-memory store for metrics
	handlers.URLStore["https://example.com"] = "abc123"
	handlers.URLStore["https://another.com"] = "def456"
	handlers.URLStore["https://example.com/test"] = "ghi789"

	req, err := http.NewRequest("GET", "/metrics", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.MetricsHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected1 := "example.com: 2"
	expected2 := "another.com: 1"
	if !contains(rr.Body.String(), expected1) || !contains(rr.Body.String(), expected2) {
		t.Errorf("handler returned unexpected body: got %v want %v or %v", rr.Body.String(), expected1, expected2)
	}
}
