package main

import (
	"fmt"
	"github.com/zilinjak/oas-proxy/internal/api"
	"github.com/zilinjak/oas-proxy/internal/config"
	"github.com/zilinjak/oas-proxy/internal/logging"
	"go.uber.org/zap"
)

func init() {
	// Load configuration
	logging.Logger.Debug("Application configuration", zap.String(
		"config", config.AppConfig.String()),
	)
}

func main() {
	// Load configuration

	// Initialize Gin router with middleware
	router := api.NewRouter()

	// Start server
	err := router.Run(":" + config.AppConfig.ServerPort)
	if err != nil {
		fmt.Printf("Error when starting the server: %s", err)
	}
}
