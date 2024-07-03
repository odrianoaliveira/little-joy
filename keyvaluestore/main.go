package main

import (
	"go.uber.org/zap"
	"keyvaluestore/middleware"
	"keyvaluestore/server"
	"log"
	"net/http"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Error creating logger: %v", err)
	}

	logger.Info("Starting server...")

	mux := http.NewServeMux()
	mux.HandleFunc("/pair", server.CreatePairHandler)
	wrappedMux := middleware.LogRequest(mux)
	logger.Info("Server is listening on port 8080...")
	err = http.ListenAndServe(":8080", wrappedMux)
	if err != nil {
		logger.Fatal("Error starting server", zap.Error(err))
	}
}
