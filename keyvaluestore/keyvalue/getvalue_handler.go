package keyvalue

import (
	"encoding/json"
	"net/http"
)

type KeyValueResponse struct {
	Key   string
	Value interface{}
}

func GetValueByKey(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid Method Request", http.StatusMethodNotAllowed)
		return
	}

	key := r.PathValue("key")

	response := KeyValueResponse{
		Key:   key,
		Value: "a value", // TODO reade from the hash map
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error serializing response", http.StatusInternalServerError)
		return
	}
}
