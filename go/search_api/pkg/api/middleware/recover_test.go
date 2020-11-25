package middleware_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ewgRa/test_tasks/go/search_api/pkg/api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Here we test how good we recover from panic inside middleware.
// This can't be a parallel test, since we replace log.Logger.
//nolint:paralleltest
func TestRecoverFromMiddlewarePanic(t *testing.T) {
	oldLogger, logWriter := mockLogger()

	defer func() { log.Logger = oldLogger }()

	panicMiddleware := func(ctx *gin.Context) {
		panic("test panic recover from middleware")
	}

	response := request([]gin.HandlerFunc{panicMiddleware}, nil)

	if response.Code != http.StatusInternalServerError {
		t.Errorf("Expected to have internal server error status code")

		return
	}

	if !strings.Contains(logWriter.String(), "test panic recover from middleware") {
		t.Errorf("Can't find log record about panic")

		return
	}
}

// Here we test how good we recover from panic inside handler.
// This can't be a parallel test, since we replace log.Logger.
//nolint:paralleltest
func TestRecoverFromHandlerPanic(t *testing.T) {
	oldLogger, logWriter := mockLogger()

	defer func() { log.Logger = oldLogger }()

	panicHandler := func(ctx *gin.Context) {
		panic("test panic recover from handler")
	}

	response := request(nil, []gin.HandlerFunc{panicHandler})

	if response.Code != http.StatusInternalServerError {
		t.Errorf("Expected to have internal server error status code")

		return
	}

	if !strings.Contains(logWriter.String(), "test panic recover from handler") {
		t.Errorf("Can't find log record about panic")

		return
	}
}

func mockLogger() (zerolog.Logger, *bytes.Buffer) {
	oldLogger := log.Logger
	logWriter := bytes.NewBufferString("")
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: logWriter})

	return oldLogger, logWriter
}

func request(middlewares []gin.HandlerFunc, handlers []gin.HandlerFunc) *httptest.ResponseRecorder {
	engine := gin.New()
	engine.Use(middleware.RecoverMiddleware)
	engine.Use(middleware.CorrelationIDMiddleware)
	engine.Use(middlewares...)
	engine.GET("/", handlers...)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	engine.HandleContext(c)

	return response
}
