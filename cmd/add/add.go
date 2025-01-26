package add

import (
	"fmt"
	"strings"
	"task/db/actions"

	"github.com/spf13/cobra"
)

var _ = fmt.Println

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		actions.Actions("add", strings.Join(args, " "))
	},
}
