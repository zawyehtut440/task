package cmd

import (
	"fmt"
	"os"

	"task/cmd/add"
	"task/cmd/completed"
	"task/cmd/do"
	"task/cmd/list"
	"task/cmd/remove"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "task is a CLI for managing your TODOs.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// hide completion and help subcommand
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true

	rootCmd.AddCommand(add.AddCmd)
	rootCmd.AddCommand(do.DoCmd)
	rootCmd.AddCommand(list.ListCmd)
	rootCmd.AddCommand(remove.RemoveCmd)
	rootCmd.AddCommand(completed.CompletedCmd)
}
