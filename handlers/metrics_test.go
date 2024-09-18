package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
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

	// More flexible comparison to check that the body contains the required metrics
	expected1 := "example.com: 2"
	expected2 := "another.com: 1"

	// Using strings.Contains instead of a strict match
	body := rr.Body.String()
	if !strings.Contains(body, expected1) || !strings.Contains(body, expected2) {
		t.Errorf("handler returned unexpected body: got %v want to contain %v and %v", body, expected1, expected2)
	}
}
