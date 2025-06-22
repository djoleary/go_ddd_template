package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	args := os.Args[1:]

	if err := run(args); err != nil {
		log.Fatalf("runtime error: %v", err)
	}
}

// run is used in place of main as it is a normal go function that can return an error.
// having explicit inputs also allows for easier testing of the whole application.
func run(args []string) error {
	cmds := []*cobra.Command{
		{
			Use:   "echo [string to echo]",
			Short: "Echo anything",
			Long:  "echo is for echoing back whatever you want",
			Args:  cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("Echo: " + strings.Join(args, " "))
			},
		},
	}

	rootCmd := &cobra.Command{Use: "cli"}
	rootCmd.SetArgs(args)
	rootCmd.AddCommand(cmds...)

	if err := rootCmd.Execute(); err != nil {
		return fmt.Errorf("cobra error: %w", err)
	}

	return nil
}
