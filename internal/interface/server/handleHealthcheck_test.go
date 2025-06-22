package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/djoleary/go_ddd_template/internal/interface/server"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestHealthcheckEndpoint(t *testing.T) {
	t.Parallel()

	mockEnv := testEnv{}
	ws := echo.New()
	s := server.NewServer(mockEnv, ws)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/healthcheck", nil)

	s.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
}
