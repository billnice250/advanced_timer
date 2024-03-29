package advanced_timer

import (
	"fmt"
	"sync"
	"time"
)

// AdvancedTimer is a wrapper around the time.Timer struct
// that is multi-thread safe and has a few extra features
//
// Specs:
// - Start() starts the timer.
// - Pause() pauses the timer.
// - Resume() resumes the timer.
// - Stop() stops the timer.
// - IsFinished() channel is closed and receives true when the timer is stopped.
// - Remaining is the remaining time on the timer.
// - MaxDuration is the maximum duration of the timer.
// - StartedAt is the time the timer was started.
// - Paused is a boolean indicating whether the timer is paused or not.
//
// Limitations:
// - The remaining time has a margin of error of around 30 milliseconds.
// - The finished event is only guaranteed to be fired within 100 microseconds of the timer being stopped.
type AdvancedTimer struct {
	inner       *time.Timer
	MaxDuration time.Duration
	StartedAt   time.Time
	Remaining   time.Duration
	finished    chan bool
	Paused      bool
	mutex       sync.Mutex
}

// NewAdvancedTimer creates a new AdvancedTimer
func NewAdvancedTimer(maxDuration time.Duration) AdvancedTimer {
	return AdvancedTimer{
		MaxDuration: maxDuration,
		finished:    make(chan bool),
	}
}

// get Finished channel read only
func (t *AdvancedTimer) IsFinished() <-chan bool {
	return t.finished
}

// Stop stops the timer and marks it as finished
func (t *AdvancedTimer) Stop() {
	t.inner.Reset(0) // This will cause the timer to fire immediately
}

// Start starts the timer with the given duration
func (t *AdvancedTimer) Start() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.Remaining = t.MaxDuration
	t.StartedAt = time.Now()
	t.inner = time.NewTimer(t.MaxDuration)
	go func() {
		<-t.inner.C
		t.mutex.Lock()
		defer t.mutex.Unlock()
		t.Remaining -= time.Since(t.StartedAt)
		t.finished <- true
	}()
}

// Pause pauses the timer and saves the remaining time
func (t *AdvancedTimer) Pause() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if !t.Paused {
		t.Paused = true
		t.Remaining -= time.Since(t.StartedAt)
		t.inner.Stop()
	}
}

// Resume resumes the timer with the remaining time
func (t *AdvancedTimer) Resume() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if t.Paused {
		t.Paused = false
		t.StartedAt = time.Now()
		t.inner.Reset(t.Remaining)
	}
}

// Stringfy returns a string representation of the AdvancedTimer
func (t *AdvancedTimer) Stringfy() string {
	return fmt.Sprintf("MaxAllowed Time: %v, Remaining: %v, Paused: %v  , Finished: %v", t.MaxDuration, t.Remaining, t.Paused, t.finished)
}
