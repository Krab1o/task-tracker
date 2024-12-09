package add

import (
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command {
    Use:    "add",
    Short:  "Adds a task",
    Long:   `Adds a task to the list of your tasks. 
			Default state is "to do".`,
    Run: func(cmd *cobra.Command, args []string) {
        addTask(args[0])
    },
}

func init() {
    
}