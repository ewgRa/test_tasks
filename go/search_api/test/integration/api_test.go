package integration_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ewgra/go-test-task/pkg/api"
	"github.com/ewgra/go-test-task/pkg/api/middleware"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

func TestAPI(t *testing.T) {
	t.Parallel()

	var cfg api.Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		t.Errorf(errors.WithMessage(err, "Can't process config").Error())

		return
	}

	s, err := api.CreateAPIEngine(&cfg)
	if err != nil {
		t.Errorf(errors.WithMessage(err, "Can't create server").Error())

		return
	}

	request := httptest.NewRequest(http.MethodGet, "/v1/products?q=Jeans", nil)
	response := httptest.NewRecorder()
	err = addAuthorization(s, request)

	if err != nil {
		t.Errorf(errors.WithMessage(err, "Can't authorize request").Error())

		return
	}

	s.ServeHTTP(response, request)

	wantStatus := 200

	if response.Code != wantStatus {
		t.Errorf("Got %v response code, %v expected", response.Code, wantStatus)

		return
	}

	got := response.Body.String()
	want := `{"data":[{"title":"512 Slim Taper Fit Jeans","brand":"LEVI'S","price":119.95,"stock":4}]}`

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

func addAuthorization(handler http.Handler, r *http.Request) error {
	requestBody := strings.NewReader(`{"username": "test", "password": "test"}`)
	request := httptest.NewRequest(http.MethodPost, "/v1/login", requestBody)
	request.Header.Set("Content-type", "application/json")

	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		return errors.New("Unexpected response code")
	}

	var data struct {
		Token string `json:"token"`
	}

	err := json.Unmarshal(response.Body.Bytes(), &data)
	if err != nil {
		return errors.WithMessage(err, "Can't unmarshal response")
	}

	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", data.Token))

	return nil
}
