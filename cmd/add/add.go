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
        // fmt.Printf("Addition of %s and %s = %s.\n\n", args[0], args[1])
    },
}

func init() {
    
}