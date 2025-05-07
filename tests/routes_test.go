//go:build !integration
// +build !integration

package api

import (
	"github.com/zilinjak/oas-proxy/internal/api"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthcheckEndpoint(t *testing.T) {
	router := api.NewRouter()
	req, _ := http.NewRequest("GET", "/oas-proxy/healthcheck", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", recorder.Code)
	}

	expected := `{"status":"ok"}`
	if recorder.Body.String() != expected+"\n" && recorder.Body.String() != expected {
		t.Errorf("Expected body %s, got %s", expected, recorder.Body.String())
	}
}
