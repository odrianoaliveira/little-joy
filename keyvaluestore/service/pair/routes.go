package pair

import (
	"encoding/json"
	"go.uber.org/zap"
	"io"
	"keyvaluestore/service"
	"net/http"
)

type Handler struct {
	logger *zap.Logger
}

func NewHandler(logger *zap.Logger) *Handler {
	return &Handler{logger: logger}
}

func (h Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/pair", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createPairHandler(h.logger, w, r)
		default:
			http.Error(w, service.MsgMethodNotAllowed, http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/pair/{key}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			setValueHandler(w, r)
		case http.MethodGet:
			getValueHandler(w, r)
		default:
			http.Error(w, service.MsgMethodNotAllowed, http.StatusMethodNotAllowed)
		}
	})
}

func createPairHandler(logger *zap.Logger, w http.ResponseWriter, r *http.Request) {
	var p Pair[string]

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, service.MsgBadRequest+": invalid JSON", http.StatusBadRequest)
		return
	}

	logger.Info("Creating pair", zap.String("key", p.Key), zap.String("value", p.Value))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	jErr := json.NewEncoder(w).Encode(p)
	if jErr != nil {
		http.Error(w, service.MsgInternalServerError, http.StatusInternalServerError)
		return
	}
}

func getValueHandler(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	p := Pair[string]{
		Key:   key,
		Value: "a value",
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		http.Error(w, service.MsgInternalServerError, http.StatusInternalServerError)
		return
	}
}

func setValueHandler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Bad Request: missing value", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	valueBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, service.MsgInternalServerError, http.StatusInternalServerError)
		return
	}
	value := string(valueBytes)
	if len(value) <= 0 {
		http.Error(w, service.MsgBadRequest+": missing value", http.StatusBadRequest)
		return
	}

	key := r.PathValue("key")
	p := Pair[string]{
		Key:   key,
		Value: value,
	}

	w.Header().Set("Content-Type", "application/json")
	jErr := json.NewEncoder(w).Encode(p)
	if jErr != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
