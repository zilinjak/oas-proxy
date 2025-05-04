package controllers

import (
	"github.com/gin-gonic/gin"
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
	oasResponse := &http.Response{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       resp.Body,
	}
	oasRequest := &http.Request{
		Method: c.Request.Method,
		URL:    c.Request.URL,
		Header: c.Request.Header,
		Body:   c.Request.Body,
	}
	proxy.OAS3Validator.Validate(oasRequest, oasResponse)
}
