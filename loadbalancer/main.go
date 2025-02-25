package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Server represents a backend server
type Server struct {
	URL          *url.URL
	Alive        bool
	ReverseProxy *httputil.ReverseProxy
}

// List of backend servers
var servers = []*Server{}

func main() {
	// Initialize backend servers
	backends := []string{
		"http://localhost:8081",
		"http://localhost:8082",
	}

	// Setup backend servers
	for _, backend := range backends {
		url, _ := url.Parse(backend)
		servers = append(servers, &Server{
			URL:          url,
			Alive:        true,
			ReverseProxy: httputil.NewSingleHostReverseProxy(url),
		})
	}

	// Create a simple round-robin load balancer
	currentServer := 0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if len(servers) > 0 {
			servers[currentServer].ReverseProxy.ServeHTTP(w, r)
			currentServer = (currentServer + 1) % len(servers)
		} else {
			http.Error(w, "No backends available", http.StatusServiceUnavailable)
		}
	})

	// Start the load balancer on port 8080
	fmt.Println("Load Balancer started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}