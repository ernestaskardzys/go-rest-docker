package httpfunctions

import (
	"log"
	"net/http"
)

func PutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Put function has been called")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}
