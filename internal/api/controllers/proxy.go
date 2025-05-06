package controllers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/zilinjak/oas-proxy/internal/api/services"
	"github.com/zilinjak/oas-proxy/internal/api/validators"
	"github.com/zilinjak/oas-proxy/internal/logging"
	"io"
	"net/http"
)

type ProxyController struct {
	ProxyService  *services.ProxyService
	OAS3Validator *validators.OAS3Validator
}

func NewProxyController(oasPath string) *ProxyController {
	return &ProxyController{
		ProxyService:  services.NewProxyService(),
		OAS3Validator: validators.NewOAS3Validator(oasPath),
	}
}

func (proxy *ProxyController) HandleTraffic(c *gin.Context) {
	// copy the request body
	// We need 2 copies of the request body: one for the OAS3 validator and one for the proxy
	oasBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		proxy.handleError(c, err.Error())
		return
	}
	forwardingBody := io.NopCloser(bytes.NewBuffer(oasBody))

	resp, err := proxy.ProxyService.Forward(c, forwardingBody)
	if err != nil {
		proxy.handleError(c, err.Error())
		return
	}
	err = proxy.ProxyService.CreateResponse(c, resp)
	if err != nil {
		proxy.handleError(c, "failed to create response: "+err.Error())
		return
	}
	proxy.OAS3Validator.Validate(c.Request, oasBody, resp)
}

func (proxy *ProxyController) handleError(c *gin.Context, message string) {
	logging.Logger.Error("Error reading request body: " + message)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})

}
