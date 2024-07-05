package keyvalue

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SetValueHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method Request", http.StatusMethodNotAllowed)
		return
	}

	key := r.PathValue("key")
	if r.Body == nil {
		http.Error(w, "Bad Request: missing value", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	valueBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	value := string(valueBytes)
	if len(value) <= 0 {
		http.Error(w, "Bad Request: missing value", http.StatusBadRequest)
		return
	}

	response := response{
		Message: fmt.Sprintf("key %s and keyvalue %s has been added.", key, value),
	}

	w.Header().Set("Content-Type", "application/json")
	jErr := json.NewEncoder(w).Encode(response)
	if jErr != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
