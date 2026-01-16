package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests, labeled by method and route",
		},
		[]string{"method", "route"},
	)
)

func init() {
	prometheus.MustRegister(requests)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Increment with labels
	requests.WithLabelValues(r.Method, r.URL.Path).Inc()
	w.Write([]byte("Hello, keda!"))
}

func main() {
	// Expose metrics endpoint
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/hello", handler)

	http.ListenAndServe(":8080", nil)
}
