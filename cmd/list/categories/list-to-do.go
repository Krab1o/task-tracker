package category

import (
	listlogic "task-tracker/internal/list_logic"
	"task-tracker/internal/status"

	"github.com/spf13/cobra"
)

var ListToDoCmd = &cobra.Command {
    Use:    "to-do",
    Short:  "Shows all tasks",
    Long:   `Shows all tasks with specified group. 
			Default state is "to do".`,
    Run: func(cmd *cobra.Command, args []string) {
        listlogic.ListTask([]string{status.ToDo.String()})
    },
}

func init() {
    
}