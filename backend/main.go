package main

import (
	"log"
	"net/http"
)

func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/health", withCORS(healthHandler))
	mux.HandleFunc("/api/add", withCORS(addHandler))
	mux.HandleFunc("/api/subtract", withCORS(subtractHandler))
	mux.HandleFunc("/api/multiply", withCORS(multiplyHandler))
	mux.HandleFunc("/api/divide", withCORS(divideHandler))
	mux.HandleFunc("/api/power", withCORS(powerHandler))
	mux.HandleFunc("/api/sqrt", withCORS(sqrtHandler))
	mux.HandleFunc("/api/percentage", withCORS(percentageHandler))

	log.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}