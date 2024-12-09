package markdone

import (
	"github.com/spf13/cobra"
)

var MarkDoneCmd = &cobra.Command {
    Use:    "mark-done",
    Short:  "Marks your task done",
    Long:   `Changes your task status to "done"`,
    Run: func(cmd *cobra.Command, args []string) {
        markDoneTask()
    },
}

func init() {
    
}