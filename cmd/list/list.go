package list

import (
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command {
    Use:    "list",
    Short:  "Shows all tasks",
    Long:   `Shows all tasks with specified group. 
			Default state is "to do".`,
    Run: func(cmd *cobra.Command, args []string) {
        listTask()
        // fmt.Printf("Addition of %s and %s = %s.\n\n", args[0], args[1])
    },
}

func init() {
    
}