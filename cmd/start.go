package cmd

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/deadpyxel/flowmodoro-cli/internal/state"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts a Flowmodoro session",
	Long:  "Starts a new Flomodoro session. The timer will run until you tell it to stop.",
	RunE:  startSession,
}

func startSession(cmd *cobra.Command, args []string) error {
	statePath := viper.GetString("statePath")
	st, err := state.LoadState(statePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			_, err := os.Create(statePath)
			if err != nil {
				return err
			}
		}
	}
	// If there is an active session, send a message and exit early
	if st.SessionActive {
		fmt.Println("A flowmodoro session is already in progress")
		return nil
	}
	// Set the current start time and the session as active
	st.StartTime = time.Now()
	st.StopTime = time.Time{} // Clear StopTime
	st.SessionActive = true
	err = state.SaveState(st, statePath)
	if err != nil {
		return err
	}
	fmt.Printf("Flowmodoro sessions started at %v\n", st.StartTime.Format("15:04:05"))
	return nil
}

func init() {
	rootCmd.AddCommand(startCmd)
}
