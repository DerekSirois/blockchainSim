package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Encode[t any](w http.ResponseWriter, statusCode int, data t) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("failed to encode the data: %v\n", err)
	}
}

func Decode[t any](r *http.Request) (t, error) {
	var data t
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return data, fmt.Errorf("failed to decode the data: %v", err)
	}
	return data, nil
}
