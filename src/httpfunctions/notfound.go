package httpfunctions

import (
	"log"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("NotFound function has been called")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}
