package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var breakCmd = &cobra.Command{
	Use:   "break",
	Short: "Starts a break timer for yout Flowmodoro sessions",
	Long:  "Starts a coutdown timer for your break session. By default the ratio of work/break time is 1/5",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("break timer started")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(breakCmd)
}
