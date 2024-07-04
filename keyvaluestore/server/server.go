package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type KeyValuePayload struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func CreatePairHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method Request", http.StatusMethodNotAllowed)
		return
	}

	var payload KeyValuePayload

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	response := map[string]string{
		"message": fmt.Sprintf("key %s and value %s has been added.", payload.Key, payload.Value),
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
