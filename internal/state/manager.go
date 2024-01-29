package state

import (
	"encoding/json"
	"os"
	"time"
)

type FlowmodoroState struct {
	SessionActive bool      `json:"session_active"`
	StartTime     time.Time `json:"start_time"`
	StopTime      time.Time `json:"stop_time"`
}

func SaveState(state FlowmodoroState, filepath string) error {
	data, err := json.Marshal(state)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath, data, 0644)
}

func LoadState(filepath string) (FlowmodoroState, error) {
	var state FlowmodoroState
	data, err := os.ReadFile(filepath)
	if err != nil {
		return FlowmodoroState{}, err
	}
	err = json.Unmarshal(data, &state)
	return state, err
}
