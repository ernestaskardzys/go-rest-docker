package main

import (
	"log"
	"net/http"
	"go-rest-docker/httpfunctions"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", httpfunctions.GetHandler).Methods(http.MethodGet)
	r.HandleFunc("/", httpfunctions.PostHandler).Methods(http.MethodPost)
	r.HandleFunc("/", httpfunctions.PutHandler).Methods(http.MethodPut)
	r.HandleFunc("/", httpfunctions.DeleteHandler).Methods(http.MethodDelete)
	r.HandleFunc("/", httpfunctions.NotFoundHandler)
	log.Println("REST service starting!")
	log.Fatal(http.ListenAndServe(":8080", r))
}