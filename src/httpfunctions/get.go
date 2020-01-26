package httpfunctions

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type GetResultDto struct {
	Name      string
	Number 	   int
	Date       time.Time
}

type GetResults []GetResultDto

func GetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Get function has been called")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	getResult := GetResults {
		GetResultDto{
			Name:   "First result",
			Number: rand.Intn(100),
			Date:   time.Time{},
		},
		GetResultDto{
			Name:   "Second result",
			Number: rand.Intn(100),
			Date:   time.Time{},
		},
	}

	json.NewEncoder(w).Encode(getResult)
}
