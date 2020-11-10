package middleware_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/ewgra/go-test-task/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func TestRecover(t *testing.T) {
	logWriter := bytes.NewBufferString("")
	oldLogger := log.Logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: logWriter})
	defer func() { log.Logger = oldLogger }()

	r := gin.New()
	r.Use(middleware.RecoverMiddleware)
	r.Use(middleware.CorrelationIDMiddleware)
	r.Use(panicMiddleware)
	r.GET("/")

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	r.HandleContext(c)

	if response.Code != http.StatusInternalServerError {
		t.Errorf("Expected to have internal server error status code")
		return
	}

	if !strings.Contains(logWriter.String(), "test panic recover") {
		t.Errorf("Can't find log record about panic")
		return
	}
}

func panicMiddleware(c *gin.Context) {
	panic(errors.New("test panic recover"))
}
