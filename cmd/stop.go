package cmd

import (
	"fmt"
	"time"

	"github.com/deadpyxel/flowmodoro-cli/internal/state"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops a Flowmodoro session",
	Long:  "Stops the current Flowmodoro session",
	RunE:  stopSession,
}

func stopSession(cmd *cobra.Command, args []string) error {
	st, err := state.LoadState("state.json")
	if err != nil {
		return err
	}
	if !st.SessionActive {
		fmt.Println("No flowmodoro session is already in progress")
		return nil
	}
	st.SessionActive = false
	st.StopTime = time.Now()

	err = state.SaveState(st, "state.json")
	if err != nil {
		return err
	}

	fmt.Printf("Flowmodoro session stopped at %v\n", st.StartTime.Format("15:04:05"))
	return nil
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
