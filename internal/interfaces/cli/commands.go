package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func (c *cli) commands() {
	cmds := []*cobra.Command{
		{
			Use:   "echo [string to echo]",
			Short: "Echo anything",
			Long:  "echo is for echoing back whatever you want",
			Args:  cobra.MinimumNArgs(1),
			RunE: func(cmd *cobra.Command, args []string) error {
				c.logger.Debug("Echoing")
				if _, err := fmt.Fprintln(cmd.OutOrStdout(), "Echo: "+strings.Join(args, " ")); err != nil {
					return err
				}
				c.logger.Debug("Echoed")
				return nil
			},
		},
	}

	c.rootCmd.AddCommand(cmds...)
}
