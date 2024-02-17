package state

import (
	"os"
	"testing"
	"time"
)

func compareState(a, b FlowmodoroState) bool {
	if a.SessionActive != b.SessionActive {
		return false
	}
	if a.StartTime != b.StartTime {
		return false
	}
	if a.StopTime != b.StopTime {
		return false
	}
	return true
}

func TestSaveState(t *testing.T) {
	filepath := "t.json"
	t.Run("When state is nil SaveState returns an error", func(t *testing.T) {
		defer os.Remove(filepath)

		st := FlowmodoroState{}
		err := SaveState(st, filepath)
		if err != nil {
			t.Errorf("SaveState(%+v, %s) error = %v", st, filepath, err)
		}
	})
}

func TestLoadState(t *testing.T) {
	t.Run("When LoadState is called with missing file it returns an error", func(t *testing.T) {
		filepath := "idontexist.json"
		_, err := LoadState(filepath)
		if err == nil {
			t.Errorf("LoadState(%s) expected an error, got %v instead", filepath, err)
		}
	})
}

func TestStateManagementCycle(t *testing.T) {
	filepath := "t.json"
	t.Run("When state is saved file contents match struct", func(t *testing.T) {
		st := FlowmodoroState{SessionActive: true, StartTime: time.Now(), StopTime: time.Now().Add(2 * time.Minute)}
		defer os.Remove(filepath)

		err := SaveState(st, filepath)
		if err != nil {
			t.Errorf("SaveState() error = %v", err)
		}

		savedSt, err := LoadState(filepath)
		if err != nil {
			t.Fatal(err)
		}
		if compareState(st, savedSt) {
			t.Errorf("Expected loaded state to match previously saved state, was %v but got %v", st, savedSt)
		}
	})
}
