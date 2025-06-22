package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func (c *cli) handleGreeting() handlerFunc {
	return func(cmd *cobra.Command, args []string) error {
		n := "world"
		if len(args) > 0 {
			n = strings.Join(args, " and ")
		}

		c.logger.Debug("Saying hello to: " + n)
		if _, err := fmt.Fprintln(cmd.OutOrStdout(), "hello "+n+"!"); err != nil {
			return err
		}
		return nil
	}
}
