package list

import (
	category "task-tracker/cmd/list/categories"
	listlogic "task-tracker/internal/list_logic"
	"task-tracker/internal/status"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command {
    Use:    "list",
    Short:  "Shows all tasks",
    Long:   `Shows all tasks with specified group. 
			Default state is "to do".`,
    Run: func(cmd *cobra.Command, args []string) {
        listlogic.ListTask(
			[]string{
				status.ToDo.String(), 
				status.InProgress.String(), 
				status.Done.String(),
			})
    },
}

func init() {
    ListCmd.AddCommand(category.ListAllCmd)
    ListCmd.AddCommand(category.ListDoneCmd)
    ListCmd.AddCommand(category.ListToDoCmd)
    ListCmd.AddCommand(category.ListInProgressCmd)
}