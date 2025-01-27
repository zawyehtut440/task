package completed

import (
	"github.com/spf13/cobra"
	"task/db/actions"
)

var CompletedCmd = &cobra.Command{
	Use:   "completed",
	Short: "List all of your complete tasks",
	Run: func(cmd *cobra.Command, args []string) {
		actions.Actions("completed")
	},
}
