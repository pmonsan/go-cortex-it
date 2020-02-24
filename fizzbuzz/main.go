package main

import (
	"fizzbuzz/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", services.Request).Methods(http.MethodGet)
	r.HandleFunc("/statistics", services.Statistics).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", r))
}
