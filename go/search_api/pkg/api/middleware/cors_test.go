package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ewgra/go-test-task/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func TestCors(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodOptions, "/", strings.NewReader(""))
	c.Request.Header.Set("Origin", "http://localhost")
	corsMiddleware := middleware.CorsMiddleware("*")
	corsMiddleware(c)
	headerValue := response.Header().Get("Access-Control-Allow-Headers")

	if !strings.Contains(headerValue, "Authorization") {
		t.Errorf("Can't find Authorization in access control allowed headers")
		return
	}
}
