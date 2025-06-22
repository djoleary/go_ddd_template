package server_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/djoleary/go_ddd_template/internal/interface/server"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

type testEnv struct {
	pairs map[string]string
}

func (e testEnv) Getenv(k string) string {
	if v, ok := e.pairs[k]; ok {
		return v
	}
	return ""
}

func TestGreetingEndpoint(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		path     string
		expected string
	}{
		"greet the world": {
			path:     "/",
			expected: "hello world!",
		},
		"greet 'ted'": {
			path:     "/ted",
			expected: "hello ted!",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			mockEnv := testEnv{}
			ws := echo.New()
			s := server.NewServer(mockEnv, ws)

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", test.path, nil)

			s.ServeHTTP(w, r)

			assert.Equal(t, http.StatusOK, w.Result().StatusCode)

			body, _ := io.ReadAll(w.Result().Body)
			assert.Equal(t, test.expected, string(body))
		})
	}
}
