package do

import (
	"strconv"
	"task/db/actions"

	"github.com/spf13/cobra"
)

var DoCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// do command accept only one argument, check that argument whether is integer
		if _, err := strconv.Atoi(args[0]); err != nil {
			return err
		}
		actions.Actions("do", args[0])
		return nil
	},
}
