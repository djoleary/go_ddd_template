package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/djoleary/go_ddd_template/internal/infrastructure/environ"
	"github.com/djoleary/go_ddd_template/internal/infrastructure/gommon"
	"github.com/djoleary/go_ddd_template/internal/infrastructure/slog"
	"github.com/djoleary/go_ddd_template/internal/interface/server"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	_ = godotenv.Load() // We don't care if this fails

	stderr := os.Stderr
	env := environ.NewOSEnv()

	if err := run(stderr, env); err != nil {
		log.Fatalf("runtime error: %v", err)
	}
}

// run is used in place of main as it is a normal go function that can return an error.
// having explicit inputs also allows for easier testing of the whole application.
func run(stderr io.Writer, env environ.Getenver) error {
	slog := slog.NewJsonLogger(stderr, env.Getenv("APP_LOG_LEVEL"))
	slog.SetAsDefault()

	ws := echo.New()
	ws.Logger.SetLevel(gommon.GetLevel(env.Getenv("APP_LOG_LEVEL")))
	ws.Logger.SetOutput(stderr)
	ws.StdLogger.SetOutput(stderr)

	s := server.NewServer(env, ws)

	if err := s.Serve(); err != http.ErrServerClosed {
		return fmt.Errorf("server error: %w", err)
	}

	return nil
}
