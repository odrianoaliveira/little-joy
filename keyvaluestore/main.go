package main

import (
	"go.uber.org/zap"
	"keyvaluestore/cmd/api"
	"log"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Error creating logger: %v", err)
	}
	logger.Info("Starting service...")

	server := api.NewAPIServer(":8080", logger)
	if err := server.Run(); err != nil {
		logger.Error("Error starting server", zap.Error(err))
	}

	logger.Info("Server is listening on port 8080...")
}
