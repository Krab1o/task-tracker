package category

import (
	listlogic "task-tracker/internal/list_logic"
	"task-tracker/internal/status"

	"github.com/spf13/cobra"
)

var ListDoneCmd = &cobra.Command {
    Use:    "done",
    Short:  "Shows all tasks",
    Long:   `Shows all tasks with specified group. 
			Default state is "to do".`,
    Run: func(cmd *cobra.Command, args []string) {
        listlogic.ListTask([]string{status.Done.String()})
    },
}

func init() {
    
}