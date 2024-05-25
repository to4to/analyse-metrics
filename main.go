package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	router := http.NewServeMux()

	router.HandleFunc("POST /data", handleData)
	http.ListenAndServe(":8080", router)
}

func handleData(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET POST OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return

	}

	b, err := io.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()
	fmt.Println(string(b))

	w.Write([]byte("Hello Cors"))

	fmt.Println(string(b))
}
