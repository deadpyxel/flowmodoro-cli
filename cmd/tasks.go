package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tasksCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage the tasks during your Flowmodoro sessions",
	Long:  "Manages tasks related to flowmodoro sessions. Addin, editing and deleting tasks is possible.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Manage tasks")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(tasksCmd)
}
