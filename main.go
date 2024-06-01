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

func main() {

	router := http.NewServeMux()

	prometheus.MustRegister(pingCounter)
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

	pingCounter.Inc()
}

// 	visitCounter.With(prometheus.Labels{
// 		"path":     data.Path,
// 		"bucketID": data.BucketID,
// 	}).Inc()
// }
// func init() {

// 	registry := prometheus.NewRegistry()

// 	// Add go runtime metrics and process collectors.
// 	registry.MustRegister(
// 		visitCounter,
// 		pingCounter,
// 	)

// }

// func handleData(w http.ResponseWriter, r *http.Request) {
// 	// Set CORS headers for all responses
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
// 	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

// 	// Handle preflight request
// 	if r.Method == "OPTIONS" {
// 		w.WriteHeader(http.StatusOK)
// 		return
// 	}

// 	// Handle POST request
// 	if r.Method == "POST" {
// 		_, err := io.ReadAll(r.Body)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		defer r.Body.Close()

// 		var data Data
// 		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
// 			panic(err)
// 		}

// 		fmt.Printf("%+v\n", data)
// 		w.Write([]byte("Hello Cors"))
// 	}
// }

// func handleData(w http.ResponseWriter, r *http.Request) {
// 	// Set CORS headers for all responses
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
// 	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

// 	// Handle preflight request
// 	if r.Method == "OPTIONS" {
// 		w.WriteHeader(http.StatusOK)
// 		return
// 	}

// 	// Handle POST request
// 	if r.Method == "POST" {
// 		b, err := io.ReadAll(r.Body)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		defer r.Body.Close()

// 		fmt.Println(string(b))
// 		w.Write([]byte("Hello Cors"))
// 	}
// }

// var visitCounter = prometheus.NewCounterVec(
// 	prometheus.CounterOpts{
// 		Name: "visits",
// 		Help: "count of the visits",
// 	},
// 	[]string{"path", "bucketID"},
// )
