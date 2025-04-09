package main

import (
	"fmt"
	"github.com/zilinjak/oas-proxy/internal/api"
	"github.com/zilinjak/oas-proxy/internal/config"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize Gin router with middleware
	router := api.NewRouter(cfg)

	// Start server
	err := router.Run(":" + cfg.ServerPort)
	if err != nil {
		fmt.Printf("Error when starting the server: %s", err)
	}
}
