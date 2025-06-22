package cli

import (
	"github.com/djoleary/go_ddd_template/internal/infrastructure/environ"
	"github.com/djoleary/go_ddd_template/internal/infrastructure/slog"
	"github.com/spf13/cobra"
)

type cli struct {
	logger  slog.Logger
	env     environ.Getenver
	rootCmd cobra.Command
}

func NewCLI(logger slog.Logger, env environ.Getenver, rootCmd cobra.Command) *cli {
	c := &cli{
		logger:  logger,
		env:     env,
		rootCmd: rootCmd,
	}
	c.commands()
	return c
}

func (c *cli) Serve() error {
	return c.rootCmd.Execute()
}
