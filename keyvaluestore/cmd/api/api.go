package api

import (
	"go.uber.org/zap"
	"keyvaluestore/middleware"
	"keyvaluestore/server"
	"net/http"
)

type APIServer struct {
	addr   string
	logger *zap.Logger
}

func NewAPIServer(addr string, logger *zap.Logger) *APIServer {
	return &APIServer{
		addr:   addr,
		logger: logger,
	}
}

func (s *APIServer) Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/key-value", server.CreatePairHandler)
	wrappedMux := middleware.LogRequest(mux)

	if servErr := http.ListenAndServe(":8080", wrappedMux); servErr != nil {
		s.logger.Fatal("Error starting server", zap.Error(servErr))
	}

	return nil
}
