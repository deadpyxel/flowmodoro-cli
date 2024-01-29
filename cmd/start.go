package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var startTime time.Time
var sessionActive bool

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts a Flowmodoro session",
	Long:  "Starts a new Flomodoro session. The timer will run until you tell it to stop.",
	RunE:  startSession,
}

func startSession(cmd *cobra.Command, args []string) error {
	if sessionActive {
		fmt.Println("A flowmodoro session is already in progress")
		return nil
	}
	startTime = time.Now()
	sessionActive = true
	fmt.Printf("Flowmodoro sessions started at %v\n", startTime.Format("15:04:05"))
	return nil
}

func init() {
	rootCmd.AddCommand(startCmd)
	// initialize sessionActive as false when the application starts
	sessionActive = false
}
