package markdone

import (
	"strconv"

	"github.com/spf13/cobra"
)

var MarkDoneCmd = &cobra.Command {
    Use:    "mark-done",
    Short:  "Marks your task done",
    Long:   `Changes your task status to "done"`,
    Run: func(cmd *cobra.Command, args []string) {
        val, _ := strconv.Atoi(args[0])
        markDoneTask(uint32(val))
    },
}

func init() {
    
}