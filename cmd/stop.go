package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops a Flowmodoro session",
	Long:  "Stops the current Flowmodoro session",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Flowmodoro session stopped")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
