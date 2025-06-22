package cli

import (
	"github.com/djoleary/go_ddd_template/internal/infrastructure/slog"
	"github.com/spf13/cobra"
)

type cli struct {
	logger  *slog.Logger
	rootCmd *cobra.Command
}

func NewCLI(logger slog.Logger, rootCmd cobra.Command) *cli {
	c := &cli{
		logger:  &logger,
		rootCmd: &rootCmd,
	}
	c.commands()
	return c
}

func (c *cli) Execute() error {
	return c.rootCmd.Execute()
}
