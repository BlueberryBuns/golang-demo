package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func ZajoncGETHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Zajonc", "true")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Zajonc!"))
}

func ZajoncPOSTHandler(w http.ResponseWriter, r *http.Request) {
	zajonc := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&zajonc)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid json"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Zajonc message has been successfully processed"))
}

func CurrentTimeGETHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.Kitchen)
	w.Write([]byte(fmt.Sprintf("the current time is %v", currentTime)))
}
