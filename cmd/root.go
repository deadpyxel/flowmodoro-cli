package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "flowmodoro",
	Short: "Flowmodoro is a CLI tool for the Flowmodoro technique",
	Long: `Flowmodoro is a CLI application to implement the Flowmodoro technique.
  A time management technique were you focus first and then define your break intervals`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

}
