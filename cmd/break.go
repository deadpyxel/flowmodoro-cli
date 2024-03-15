package cmd

import (
	"fmt"

	"github.com/deadpyxel/flowmodoro-cli/internal/state"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var breakCmd = &cobra.Command{
	Use:   "break",
	Short: "Starts a break timer for your Flowmodoro sessions",
	Long:  "Starts a countdown timer for your break session. By default the ratio of work/break time is 1/5",
	RunE:  breakStart,
}

// breakStart calculates and prints the estimated break time based on the last focus session.
//
// Parameters:
// - cmd: a pointer to a cobra.Command object representing the command being executed
// - args: a slice of strings representing the arguments passed to the command
//
// Returns:
// - error: an error object if any error occurs during the calculation or printing of the break time
//
// This function loads the state from the specified statePath using the LoadState function from the state package.
// It then checks if there is an active session in progress and if there is a stop time defined for the last session.
// If there is an active session, it prints a message instructing the user to stop the session first.
// If there is no stop time defined for the last session, it returns an error indicating that there is no stop time defined.
// It then calculates the break length as 1/5 of the duration of the last focus session and prints the estimated break time.
// If the stop time is before the start time, it returns an error indicating that the stop time cannot happen before the start time.
func breakStart(cmd *cobra.Command, args []string) error {
	statePath := viper.GetString("statePath")
	st, err := state.LoadState(statePath)
	if err != nil {
		return err
	}
	if st.SessionActive {
		fmt.Println("There is a flowmodoro session is already in progress. Stop it first with the `stop` command")
		return nil
	}
	if st.StopTime.IsZero() {
		return fmt.Errorf("There's no Stop time defined for the last session")
	}
	// Calculate break length, 1/5 of the last focus session
	breakLength := st.StopTime.Sub(st.StartTime) / 5
	if st.StopTime.Before(st.StartTime) {
		return fmt.Errorf("Stop cannot happen before start")
	}
	fmt.Printf("Estimated break time: %v", breakLength)

	return nil
}

func init() {
	rootCmd.AddCommand(breakCmd)
}
