package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
)

var muMetrics sync.Mutex // Mutex for metrics

// MetricsHandler returns the top 3 domains that have been shortened the most
func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	domainCount := make(map[string]int)

	muMetrics.Lock()
	for originalURL := range URLStore {
		u, _ := url.Parse(originalURL)
		domain := strings.Split(u.Hostname(), ":")[0]
		domainCount[domain]++
	}
	muMetrics.Unlock()

	type domainStat struct {
		domain string
		count  int
	}
	var stats []domainStat
	for domain, count := range domainCount {
		stats = append(stats, domainStat{domain, count})
	}

	// Sort the slice by the count in descending order
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].count > stats[j].count
	})

	// Handle case where there are fewer than 3 domains
	limit := 3
	if len(stats) < 3 {
		limit = len(stats)
	}

	// Return the top domains
	for i := 0; i < limit; i++ {
		fmt.Fprintf(w, "%s: %d\n", stats[i].domain, stats[i].count)
	}
}
