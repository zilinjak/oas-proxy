package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zilinjak/oas-proxy/internal/config"
	"github.com/zilinjak/oas-proxy/internal/logging"
	"io"
	"log"
	"net/http"
)

type ProxyController struct {
	Client *http.Client
}

func NewProxyController() *ProxyController {
	return &ProxyController{
		Client: &http.Client{},
	}
}

func (proxy *ProxyController) HandleTraffic(c *gin.Context) {
	// Clone original request properly
	headers := c.Request.Header.Clone()
	headers.Del("Host") // Remove Host header to avoid conflicts

	// Get data of the request
	data := c.Request.Body

	// Prepare URL
	reqUrl := config.AppConfig.TargetURL.Scheme + "://" + config.AppConfig.TargetURL.Host + c.Request.URL.Path

	req, err := http.NewRequest(c.Request.Method, reqUrl, data)

	if err != nil {
		logging.Logger.Error("Error creating request: " + err.Error())
		c.String(http.StatusBadRequest, "Invalid request: %v", err)
		return
	}
	req.Header = headers
	logging.Logger.Debug("Proxying request to " + req.URL.String())
	resp, err := proxy.Client.Do(req)
	if err != nil {
		c.String(http.StatusBadGateway, "Proxy error: %v", err)
		return
	}
	defer resp.Body.Close()

	logging.Logger.Info("Response status: " + http.StatusText(resp.StatusCode))

	// Copy headers
	copyHeaders(c.Writer.Header(), resp.Header)

	// Write status
	c.Status(resp.StatusCode)

	// Copy body
	_, err = io.Copy(c.Writer, resp.Body)
	if err != nil {
		log.Printf("Error copying response body: %v", err)
	}
}

// Helper function to copy headers
func copyHeaders(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
