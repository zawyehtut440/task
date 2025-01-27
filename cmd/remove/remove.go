package remove

import (
	"strconv"
	"task/db/actions"

	"github.com/spf13/cobra"
)

var RemoveCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete a specific task on your TODO list",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// do command accept only one argument, check that argument whether is integer
		if _, err := strconv.Atoi(args[0]); err != nil {
			return err
		}
		actions.Actions("rm", args[0])
		return nil
	},
}
