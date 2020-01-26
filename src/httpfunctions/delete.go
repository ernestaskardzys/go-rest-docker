package httpfunctions

import (
	"log"
	"net/http"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete function has been called")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}
