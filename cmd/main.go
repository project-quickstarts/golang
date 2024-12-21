package main

import (
	"log"
	"net/http"

	"<%= params.namespace ? `${params.namespace}/${projectName}` : projectName %>/internal/config"
	"<%= params.namespace ? `${params.namespace}/${projectName}` : projectName %>/internal/handlers"
	"<%= params.namespace ? `${params.namespace}/${projectName}` : projectName %>/internal/middlewares"
)

func main() {
	// Load the environment variables
	config.LoadEnv()

	// Get config
	cfg := config.GetConfig()

	// Register router
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.HealthCheck)
	mux.Handle("/forward",
		middlewares.ChainMiddleware(
			http.HandlerFunc(handlers.Forward),
			middlewares.LoggingMiddleware,
			middlewares.CorsMiddleware,
		))

	// Start the server
	log.Printf("Starting server on port %s\n", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
