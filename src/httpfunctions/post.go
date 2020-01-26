package httpfunctions

import (
	"log"
	"net/http"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Post function has been called")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}
