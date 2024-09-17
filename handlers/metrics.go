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
	for originalURL := range urlStore {
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

	sort.Slice(stats, func(i, j int) bool {
		return stats[i].count > stats[j].count
	})

	// Return the top 3 domains
	top3 := stats[:3]
	for _, stat := range top3 {
		fmt.Fprintf(w, "%s: %d\n", stat.domain, stat.count)
	}
}
