package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (c *cli) handleHealthcheck() handlerFunc {
	return func(cmd *cobra.Command, args []string) error {
		if _, err := fmt.Fprintln(cmd.OutOrStdout(), "healthy"); err != nil {
			return err
		}
		return nil
	}
}
