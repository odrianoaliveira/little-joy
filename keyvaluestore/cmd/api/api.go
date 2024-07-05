package api

import (
	"go.uber.org/zap"
	"keyvaluestore/keyvalue"
	"keyvaluestore/middleware"
	"net/http"
)

type Server struct {
	addr   string
	logger *zap.Logger
}

func NewServer(addr string, logger *zap.Logger) *Server {
	return &Server{
		addr:   addr,
		logger: logger,
	}
}

func (s *Server) Run() error {
	api := http.NewServeMux()
	api.HandleFunc("/key-value/{key}", keyvalue.SetValueHandler)
	wrappedMux := middleware.LogRequest(api)

	root := http.NewServeMux()
	root.Handle("/api/v1", http.StripPrefix("/api/v1", wrappedMux))

	if servErr := http.ListenAndServe(":8080", root); servErr != nil {
		s.logger.Fatal("Error starting server", zap.Error(servErr))
	}

	return nil
}
