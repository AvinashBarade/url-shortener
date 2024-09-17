package main

import (
	"log"
	"net/http"

	"github.com/AvinashBarade/url-shortener/handlers"
)

func main() {
	// Define the routes
	http.HandleFunc("/shorten", handlers.ShortenHandler) // POST /shorten -> Shortens the URL
	http.HandleFunc("/metrics", handlers.MetricsHandler) // GET /metrics -> Shows top 3 domains
	http.HandleFunc("/", handlers.RedirectHandler)       // GET /{shortURL} -> Redirect to original URL

	// Start the server on port 8080
	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
