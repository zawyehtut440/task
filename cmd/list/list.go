package list

import (
	"github.com/spf13/cobra"
	"task/db/actions"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your imcomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {
		actions.Actions("list")
	},
}
