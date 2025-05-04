package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zilinjak/oas-proxy/internal/api/services"
	"github.com/zilinjak/oas-proxy/internal/logging"
	"net/http"
)

type ProxyController struct {
	ProxyService *services.ProxyService
}

func NewProxyController() *ProxyController {
	return &ProxyController{
		ProxyService: services.NewProxyService(),
	}
}

func (proxy *ProxyController) HandleTraffic(c *gin.Context) {
	resp, err := proxy.ProxyService.Forward(c)
	// Forwarding failed
	if err != nil {
		logging.Logger.Error("Error forwarding request: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	err = proxy.ProxyService.CreateResponse(c, resp)
	// Creating response failed
	if err != nil {
		logging.Logger.Error("Error creating response: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	// TODO: Validate response against OAS
	// TODO: Validate request against OAS
}
