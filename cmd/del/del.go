package del

import (
	"github.com/spf13/cobra"

	"strconv"
)

var DelCmd = &cobra.Command {
    Use:    "del",
    Short:  "Deletes a task",
    Long:   `Deletes a task from the list of your tasks. 
			Deletes it by id.`,
    Run: func(cmd *cobra.Command, args []string) {
        val, _ := strconv.Atoi(args[0])
		delTask(uint32(val))
    },
}

func init() {
    
}