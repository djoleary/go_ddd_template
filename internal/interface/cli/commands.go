package cli

import (
	"github.com/spf13/cobra"
)

type handlerFunc func(cmd *cobra.Command, args []string) error

func (c *cli) commands() {
	cmds := []*cobra.Command{
		{
			Use:   "greet [name to greet]...",
			Short: "Says hello",
			Long:  "Says hello to the everyone that is specified",
			RunE:  c.handleGreeting(),
		},
		{
			Use:   "healthcheck",
			Short: "returns 'healthy' when called",
			Long:  "returns a simple string when called so that healthy of the application can be tested",
			RunE:  c.handleHealthcheck(),
		},
	}

	c.rootCmd.AddCommand(cmds...)
}
