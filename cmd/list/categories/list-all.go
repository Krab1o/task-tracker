package category

import (
	listlogic "task-tracker/internal/list_logic"
	"task-tracker/internal/status"

	"github.com/spf13/cobra"
)

var ListAllCmd = &cobra.Command {
    Use:    "all",
    Short:  "Shows all tasks",
    Long:   `Shows all tasks from 3 different category. 
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
    
}