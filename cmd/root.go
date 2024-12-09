package cmd

import (
	"fmt"
	"os"
	"task-tracker/cmd/add"
	"task-tracker/cmd/list"
	markdone "task-tracker/cmd/mark-done"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
    Use:   "task-tracker",
    Short: "task-tracker is a cli tool for tracking your tasks",
    Long:  "task-tracker is a cli tool for tracking your tasks. some more blah-blah",
    Run: func(cmd *cobra.Command, args []string) {
		
    },
}

func init() {
	rootCmd.AddCommand(add.AddCmd)
	rootCmd.AddCommand(list.ListCmd)
    rootCmd.AddCommand(markdone.MarkDoneCmd)
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Oops. An error while executing task-tracker '%s'\n", err)
        os.Exit(1)
    }
}