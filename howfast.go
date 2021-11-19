package howfast

import (
	"errors"
	"time"
)

// Timer is a struct that can be used to time how long your code takes to execute.
type Timer struct {
	start time.Time
	stop  time.Time
}

// Start sets the Timer's start field to Time.Now()
func (t *Timer) Start() error {
	if !t.start.IsZero() {
		return errors.New("can't start a timer that's already started, call Reset() if you want to restart the timer")
	}
	t.start = time.Now()
	return nil
}

// Stop sets the Timer's stop field to Time.Now()
func (t *Timer) Stop() error {
	if t.start.IsZero() {
		return errors.New("can't stop a timer that hasn't been started")
	}
	t.stop = time.Now()
	return nil
}

// Clear resets a Timer's start and stop fields.
func (t *Timer) Clear() {
	t.start = time.Time{}
	t.stop = time.Time{}
}

// SecondsElapsed returns the number of seconds the timer ran for.
func (t *Timer) SecondsElapsed() (int64, error) {
	if !t.start.IsZero() && t.stop.IsZero() {
		return 0, errors.New("timer is still running")
	}
	duration := t.stop.Unix() - t.start.Unix()
	return duration, nil
}

// MillisecondsElapsed returns the number of milliseconds the timer ran for.
// A millisecond is 1/1000 of a second.
func (t *Timer) MillisecondsElapsed() (int64, error) {
	if !t.start.IsZero() && t.stop.IsZero() {
		return 0, errors.New("timer is still running")
	}
	duration := t.stop.UnixMilli() - t.start.UnixMilli()
	return duration, nil
}

// MicrosecondsElapsed returns the number of milliseconds the timer ran for.
// A microsecond is 1/1000 of a millisecond.
func (t *Timer) MicrosecondsElapsed() (int64, error) {
	if !t.start.IsZero() && t.stop.IsZero() {
		return 0, errors.New("timer is still running")
	}
	duration := t.stop.UnixMicro() - t.start.UnixMicro()
	return duration, nil
}

// NanosecondsElapsed returns the number of milliseconds the timer ran for.
// A nanosecond is 1/1000 of a microsecond.
func (t *Timer) NanosecondsElapsed() (int64, error) {
	if !t.start.IsZero() && t.stop.IsZero() {
		return 0, errors.New("timer is still running")
	}
	duration := t.stop.UnixNano() - t.start.UnixNano()
	return duration, nil
}
