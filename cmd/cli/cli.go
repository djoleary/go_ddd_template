package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/djoleary/go_ddd_template/internal/infrastructure/env"
	"github.com/djoleary/go_ddd_template/internal/infrastructure/slog"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	stdin := os.Stdin
	stdout := os.Stdout
	stderr := os.Stderr
	args := os.Args[1:]
	env := env.NewOSEnv()

	if err := run(stdin, stdout, stderr, env, args); err != nil {
		log.Fatalf("runtime error: %v", err)
	}
}

// run is used in place of main as it is a normal go function that can return an error.
// having explicit inputs also allows for easier testing of the whole application.
func run(stdin io.Reader, stdout, stderr io.Writer, env env.EnvReaderInterface, args []string) error {
	slog := slog.NewJsonLogger(stderr, env.Getenv("APP_LOG_LEVEL"))

	cmds := []*cobra.Command{
		{
			Use:   "echo [string to echo]",
			Short: "Echo anything",
			Long:  "echo is for echoing back whatever you want",
			Args:  cobra.MinimumNArgs(1),
			RunE: func(cmd *cobra.Command, args []string) error {
				slog.Debug("Echoing")
				if _, err := fmt.Fprintln(stdout, "Echo: "+strings.Join(args, " ")); err != nil {
					return err
				}
				slog.Debug("Echoed")
				return nil
			},
		},
	}

	rootCmd := &cobra.Command{Use: "cli"}
	rootCmd.SetIn(stdin)
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)
	rootCmd.SetArgs(args)
	rootCmd.AddCommand(cmds...)

	if err := rootCmd.Execute(); err != nil {
		return fmt.Errorf("cobra error: %w", err)
	}

	return nil
}
