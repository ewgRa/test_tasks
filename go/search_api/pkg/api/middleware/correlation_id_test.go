package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ewgRa/test_tasks/go/search_api/pkg/api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func TestCorrelationIdGeneration(t *testing.T) {
	t.Parallel()

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	middleware.CorrelationIDMiddleware(c)

	headerValue := response.Header().Get(middleware.CorrelationIDHeader)

	if headerValue == "" {
		t.Errorf("Can't find correlation id in response header")

		return
	}

	contextValue, exists := c.Get(middleware.CorrelationIDCtxValue)

	if !exists {
		t.Errorf("Can't find correlation id in context")

		return
	}

	if contextValue != headerValue {
		t.Errorf("Correlation id in context and in header should be the same")

		return
	}

	if log.Ctx(c.Request.Context()).GetLevel() == zerolog.Disabled {
		t.Errorf("Don't see that logger for request context is configured properly")

		return
	}
}

func TestCorrelationIdProvidedInRequest(t *testing.T) {
	t.Parallel()

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	wantID := "baobab"
	c.Request, _ = http.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	c.Request.Header.Set(middleware.CorrelationIDHeader, wantID)
	middleware.CorrelationIDMiddleware(c)

	headerValue := response.Header().Get(middleware.CorrelationIDHeader)

	if headerValue != wantID {
		t.Errorf("Correlation id header should be the same as in request")

		return
	}
}
