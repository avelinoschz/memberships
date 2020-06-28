package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/alive", HealthCheck).Methods("GET")

	log.Println("Server listening on port :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// HealthCheck ...
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("HealthCheck")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("{\"alive\": true}"))
}
