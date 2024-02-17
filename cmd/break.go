package cmd

import (
	"fmt"

	"github.com/deadpyxel/flowmodoro-cli/internal/state"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var breakCmd = &cobra.Command{
	Use:   "break",
	Short: "Starts a break timer for yout Flowmodoro sessions",
	Long:  "Starts a coutdown timer for your break session. By default the ratio of work/break time is 1/5",
	RunE:  breakStart,
}

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
