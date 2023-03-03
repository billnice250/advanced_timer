package tests

import (
	"testing"
	"time"

	"github.com/billnice250/advanced_timer"
)

func TestAdvancedTimer_Start(t *testing.T) {
	allowedExecutionTime := 5 * time.Second
	timer := advanced_timer.NewAdvancedTimer(allowedExecutionTime)
	timer.Start()
}

func TestAdvancedTimer_Pause(t *testing.T) {
	allowedExecutionTime := 5 * time.Second
	timer := advanced_timer.NewAdvancedTimer(allowedExecutionTime)
	timer.Start()
	timer.Pause()

	if !timer.Paused {
		t.Errorf("expected timer state to be PAUSED")
	}
}

func TestAdvancedTimer_Stop(t *testing.T) {
	allowedExecutionTime := 5 * time.Second
	timer := advanced_timer.NewAdvancedTimer(allowedExecutionTime)

	timer.Start()
	timer.Stop()

	time.Sleep(100 * time.Microsecond)

	select {
	case <-timer.IsFinished():
	default:
		t.Errorf("expected timer state to be STOPPED but finished channel is not closed")
	}
}

func TestAdvancedTimer_Resume(t *testing.T) {
	allowedExecutionTime := 5 * time.Second
	timer := advanced_timer.NewAdvancedTimer(allowedExecutionTime)

	timer.Start()
	timer.Pause()
	timer.Resume()

	if timer.Paused {
		t.Errorf("expected timer state to be RUNNING")
	}
}

func TestAdvancedTimer_RemainingTime(t *testing.T) {
	allowedExecutionTime := 5 * time.Second
	timer := advanced_timer.NewAdvancedTimer(allowedExecutionTime)

	timer.Start()
	time.Sleep(2 * time.Second)
	timer.Pause()

	remaining := timer.Remaining.Round(time.Millisecond)
	expectedRemaining := (3 * time.Second).Round(time.Millisecond)
	margin := 100 * time.Millisecond

	if remaining < expectedRemaining-margin || remaining > expectedRemaining+margin {
		t.Errorf("expected remaining time to be around %v, but got %v", expectedRemaining, remaining)
	}
}
func TestAdvancedTimer_RemainingTime_AfterResume(t *testing.T) {
	allowedExecutionTime := 60 * time.Second
	timer := advanced_timer.NewAdvancedTimer(allowedExecutionTime)

	timer.Start()
	time.Sleep(3 * time.Second)
	timer.Pause()
	timer.Resume()
	time.Sleep(3 * time.Second)
	timer.Stop()

	// sleep for a bit to make sure the timer has stopped
	// to pass the test, the timer should have stopped before the sleep and therefore under 1ns
	time.Sleep(100 * time.Microsecond)
	select {
	case <-timer.IsFinished():
		remaining := timer.Remaining.Round(time.Millisecond)
		expectedRemaining := (allowedExecutionTime - (3+3)*time.Second).Round(time.Millisecond)
		margin := 30 * time.Millisecond

		if remaining < expectedRemaining-margin || remaining > expectedRemaining+margin {
			t.Errorf("Expected remaining time to be around %v, but got %v", expectedRemaining, remaining)
		} else {
			t.Logf("Expected remaining time to be around %v, and got %v", expectedRemaining, remaining)
		}
	default:
		t.Errorf("timer took more than 10 microseconds to send finished signal")
	}
}
