package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", wrapped(handler))

	log.Println("server started")
	log.Fatal(http.ListenAndServe(":9000", mux))
}

func wrapped(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("requested %v", r.URL)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		handler(w, r)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Code    int       `json:"code"`
		Message string    `json:"message"`
		Created time.Time `json:"created"`
	}
	enc := json.NewEncoder(w)

	err := enc.Encode(&response{1, "hello", time.Now()})

	if err != nil {
		log.Fatalf("error encoding message: %v", err)
	}

}
