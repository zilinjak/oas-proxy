package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zilinjak/oas-proxy/internal"
	"github.com/zilinjak/oas-proxy/internal/api/services"
	"github.com/zilinjak/oas-proxy/internal/api/validators"
	"github.com/zilinjak/oas-proxy/internal/logging"
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
	forwardData, oasData, err := internal.CopyData(c.Request.Body)
	c.Request.Body = forwardData
	resp, err := proxy.ProxyService.Forward(c)
	if err != nil {
		proxy.handleError(c, err.Error())
		return
	}
	responseData, oasResponseData, err := internal.CopyData(resp.Body)
	resp.Body = responseData
	err = proxy.ProxyService.CreateResponse(c, resp)
	if err != nil {
		proxy.handleError(c, "failed to create response: "+err.Error())
		return
	}
	c.Request.Body = oasData
	resp.Body = oasResponseData
	proxy.OAS3Validator.Validate(c.Request, resp)
}

func (proxy *ProxyController) handleError(c *gin.Context, message string) {
	logging.Logger.Error("Error reading request body: " + message)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})

}
