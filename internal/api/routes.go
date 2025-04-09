package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zilinjak/oas-proxy/internal/api/controllers"
	"github.com/zilinjak/oas-proxy/internal/config"
)

func NewRouter(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	// Proxy setup
	proxyController := controllers.ProxyController{}

	router.Any("/", proxyController.HandleTraffic)

	router.GET("/oas-proxy/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return router
}
