package main

import (
	"encoding/json"
	"fmt"
	
	"net/http"
)

func main() {

	router := http.NewServeMux()

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

	// Handle POST request
	if r.Method == "POST" {
		var data Data
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		fmt.Printf("%+v\n", data)
		w.Write([]byte("Hello Cors"))
	}
}

type Data struct {
	Path     string
	BucketID string
}

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
