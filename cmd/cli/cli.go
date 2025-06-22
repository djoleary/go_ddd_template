package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/djoleary/go_ddd_template/internal/infrastructure/environ"
	"github.com/djoleary/go_ddd_template/internal/infrastructure/slog"
	"github.com/djoleary/go_ddd_template/internal/interface/cli"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func main() {
	_ = godotenv.Load() // We don't care if this fails

	stdin := os.Stdin
	stdout := os.Stdout
	stderr := os.Stderr
	args := os.Args[1:]
	env := environ.NewOSEnv()

	if err := run(stdin, stdout, stderr, env, args); err != nil {
		log.Fatalf("runtime error: %v", err)
	}
}

// run is used in place of main as it is a normal go function that can return an error.
// having explicit inputs also allows for easier testing of the whole application.
func run(stdin io.Reader, stdout, stderr io.Writer, env environ.Getenver, args []string) error {
	slog := slog.NewJsonLogger(stderr, env.Getenv("APP_LOG_LEVEL"))

	rootCmd := &cobra.Command{Use: "cli"}
	rootCmd.SetIn(stdin)
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)
	rootCmd.SetArgs(args)

	c := cli.NewCLI(*slog, env, *rootCmd)

	if err := c.Serve(); err != nil {
		return fmt.Errorf("cli error: %w", err)
	}

	return nil
}
