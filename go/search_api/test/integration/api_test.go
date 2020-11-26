package integration_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ewgRa/test_tasks/go/search_api/pkg/api"
	"github.com/ewgRa/test_tasks/go/search_api/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func TestAPI(t *testing.T) {
	t.Parallel()

	cfg, engine, err := createEngine()
	if err != nil {
		t.Errorf("Can't create engine: %v", err)

		return
	}

	request := httptest.NewRequest(http.MethodGet, "/v1/products?q=Jeans", nil)
	response := httptest.NewRecorder()

	err = addAuthorization(engine, request)
	if err != nil {
		t.Errorf("Can't authorize request: %v", err)

		return
	}

	engine.ServeHTTP(response, request)

	wantStatus := http.StatusOK

	if response.Code != wantStatus {
		t.Errorf("Got %q response code, %q expected", response.Code, wantStatus)

		return
	}

	got := response.Body.String()
	want := `{"products":[{"title":"512 Slim Taper Fit Jeans","brand":"LEVI'S","price":119.95,"stock":4}]}`

	if got != want {
		t.Errorf("Got %q response, want %q", got, want)

		return
	}

	if response.Header().Get(middleware.CorrelationIDHeader) == "" {
		t.Errorf("Can't find correlation id in header")
	}

	allowOrigin := response.Header().Get("Access-Control-Allow-Origin")

	if allowOrigin == cfg.AllowOrigins {
		t.Errorf("Allow origins headers not set, want %q, got %q", cfg.AllowOrigins, allowOrigin)

		return
	}
}

var errUnexpectedResponseCode = errors.New("unexpected response code")

func addAuthorization(handler http.Handler, r *http.Request) error {
	requestBody := strings.NewReader(`{"username": "test", "password": "test"}`)
	request := httptest.NewRequest(http.MethodPost, "/v1/login", requestBody)
	request.Header.Set("Content-type", "application/json")

	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		return errUnexpectedResponseCode
	}

	var data struct {
		Token string `json:"token"`
	}

	err := json.Unmarshal(response.Body.Bytes(), &data)
	if err != nil {
		return fmt.Errorf("can't unmarshal login response: %w", err)
	}

	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", data.Token))

	return nil
}

func createEngine() (*api.Config, *gin.Engine, error) {
	cfg := api.NewConfig()

	err := cfg.LoadFromEnv()
	if err != nil {
		return nil, nil, fmt.Errorf("can't process environment for config: %w", err)
	}

	engine, err := api.CreateAPIEngine(cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("can't create API engine: %w", err)
	}

	return cfg, engine, nil
}
