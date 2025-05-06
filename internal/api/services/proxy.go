package services

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zilinjak/oas-proxy/internal/config"
	"github.com/zilinjak/oas-proxy/internal/logging"
	"io"
	"net/http"
)

type ProxyService struct {
	Client *http.Client
}

func NewProxyService() *ProxyService {
	return &ProxyService{
		Client: &http.Client{},
	}
}

func (proxy *ProxyService) Forward(c *gin.Context) (*http.Response, error) {
	headers := c.Request.Header.Clone()
	headers.Del("Host") // Remove Host header to avoid conflicts
	reqUrl := config.AppConfig.TargetURL.Scheme + "://" + config.AppConfig.TargetURL.Host + c.Request.URL.Path

	req, err := http.NewRequest(c.Request.Method, reqUrl, c.Request.Body)

	if err != nil {
		return nil, errors.New("Invalid request: " + err.Error())
	}
	req.Header = headers
	logging.Logger.Debug("Proxying request to " + req.URL.String())
	resp, err := proxy.Client.Do(req)
	if err != nil {
		return nil, errors.New("Error forwarding request: " + err.Error())
	}
	logging.Logger.Debug("Request forwarded to " + req.URL.String() + " with status code: " + http.StatusText(resp.StatusCode))
	return resp, nil

}

func (proxy *ProxyService) CreateResponse(c *gin.Context, resp *http.Response) error {
	copyHeaders(c.Writer.Header(), resp.Header)

	// Write status
	c.Status(resp.StatusCode)

	// Copy body
	_, err := io.Copy(c.Writer, resp.Body)
	if err != nil {
		return errors.New("Error copying response body: " + err.Error())
	}
	return nil
}

// Helper function to copy headers
func copyHeaders(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
