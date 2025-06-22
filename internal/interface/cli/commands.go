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
	}

	c.rootCmd.AddCommand(cmds...)
}
