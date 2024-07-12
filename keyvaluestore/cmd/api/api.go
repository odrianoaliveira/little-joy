package main

import (
	"go.uber.org/zap"
	"keyvaluestore/middleware"
	"keyvaluestore/service/pair"
	"log"
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
	wrappedMux := middleware.LogRequest(api)

	root := http.NewServeMux()
	root.Handle("/api/v1", http.StripPrefix("/api/v1", wrappedMux))
	pair.NewHandler(s.logger).RegisterRoutes(api)

	if servErr := http.ListenAndServe(":8080", root); servErr != nil {
		s.logger.Fatal("Error starting server", zap.Error(servErr))
	}

	return nil
}

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Error creating logger: %v", err)
	}
	logger.Info("Starting service...")

	server := NewServer(":8080", logger)
	if err := server.Run(); err != nil {
		logger.Error("Error starting server", zap.Error(err))
	}

	logger.Info("Server is listening on port 8080...")
}
