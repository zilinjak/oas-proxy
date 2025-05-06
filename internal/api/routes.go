package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zilinjak/oas-proxy/internal/api/controllers"
	"github.com/zilinjak/oas-proxy/internal/config"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// Proxy setup
	proxyController := controllers.NewProxyController(config.AppConfig.OASPath)

	oas := router.Group("/oas-proxy")
	{
		oas.GET("/healthcheck", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})
	}

	// Catch-all AFTER /oas-proxy is defined
	router.NoRoute(proxyController.HandleTraffic)

	return router
}
