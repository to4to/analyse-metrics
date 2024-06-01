package main

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Data struct {
	Path     string
	BucketID string
}

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "site_visits",
		Help: "count of the visits",
	},
)

var visitCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "visit_counter",
		Help: "count of the visits",
	},
	[]string{"path", "bucketID"},
)

func main() {

	router := http.NewServeMux()

	prometheus.MustRegister(pingCounter, visitCounter)
	router.Handle("/metrics", promhttp.Handler())
	router.HandleFunc("/data", handleData)
	http.ListenAndServe(":8080", router)
}

func handleData(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for all responses
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// Handle preflight request
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	var data Data
	// Handle POST request
	if r.Method == "POST" {

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		fmt.Printf("%+v\n", data)
		w.Write([]byte("Hello Cors"))

	}

	visitCounter.With(
		prometheus.Labels{
			"path":     data.Path,
			"bucketID": data.BucketID,
		},
	)
	pingCounter.Inc()
}
