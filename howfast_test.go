package howfast

import (
	"testing"
	"time"
)

func TestCreateTimer(t *testing.T) {
	timer := Timer{}
	if !timer.start.IsZero() {
		t.Fatalf("expected time.start to be initialized to the zero time, instead it was: %v", timer.start)
	}
	if !timer.start.IsZero() {
		t.Fatalf("expected time.stop to be initialized to the zero time, but instead it was: %v", timer.stop)
	}
}

func TestTimerStart(t *testing.T) {
	timer := Timer{}
	timer.Start()
	if timer.start.IsZero() {
		t.Fatalf("expected timer.start to have a non-zero value, instead it was : %v", timer.start)
	}
	err := timer.Start()
	if err == nil {
		t.Fatalf("expected calling timer.Start() on a timer that's already started to throw an error.")
	}
}

func TestTimerStop(t *testing.T) {
	timer := Timer{}
	err := timer.Stop()
	if err == nil {
		t.Fatalf("expected a call to timer.Stop() before calling timer.Start() to throw an error")
	}
	timer.Start()
	err = timer.Stop()
	if err != nil {
		t.Fatalf("no error should be thrown when calling timer.Stop() on a timer that was started.")
	}
}

func TestTimerClear(t *testing.T) {
	timer := Timer{}
	timer.Start()
	timer.Stop()
	timer.Clear()
	if !timer.start.IsZero() {
		t.Fatalf("expected timer.start to have a zero value after timer.Clear()")
	}
	if !timer.stop.IsZero() {
		t.Fatalf("expected timer.stop to have a zero value after timer.Clear()")
	}
}

func TestSecondsElapsed(t *testing.T) {
	timer := Timer{}
	timer.Start()
	// sleep for 1 second so that duration is at least 1 second
	time.Sleep(time.Second)
	timer.Stop()

	duration, _ := timer.SecondsElapsed()
	if !(duration >= 1 && duration < 2) {
		t.Fatalf("expected duration to be greater than or equal to 1 but less than 2, instead it was %v", duration)
	}
	timer2 := Timer{}
	timer2.Start()
	_, err := timer2.SecondsElapsed()
	if err == nil {
		t.Fatalf("expected an error to be thrown when calling SecondsElapsed on a timer that hasn't stopped")
	}
}

func TestMillisecondsElapsed(t *testing.T) {
	timer := Timer{}
	timer.Start()
	// sleep for 1 second so that duration is at least 1000 milliseconds
	time.Sleep(time.Second)
	timer.Stop()

	duration, _ := timer.MillisecondsElapsed()
	if !(duration >= 1000 && duration < 2000) {
		t.Fatalf("expected duration to be greater than or equal to 1000 but less than 2000, instead it was %v", duration)
	}
	timer2 := Timer{}
	timer2.Start()
	_, err := timer2.MillisecondsElapsed()
	if err == nil {
		t.Fatalf("expected an error to be thrown when calling MillisecondsElapsed on a timer that hasn't stopped")
	}
}

func TestMicrosecondsElapsed(t *testing.T) {
	timer := Timer{}
	timer.Start()
	// sleep for 1 second so that duration is at least 1000000 microseconds
	time.Sleep(time.Second)
	timer.Stop()

	duration, _ := timer.MicrosecondsElapsed()
	if !(duration >= 1000000 && duration < 2000000) {
		t.Fatalf("expected duration to be greater than or equal to 1,000,000 but less than 2,000,000, instead it was %v", duration)
	}
	timer2 := Timer{}
	timer2.Start()
	_, err := timer2.MicrosecondsElapsed()
	if err == nil {
		t.Fatalf("expected an error to be thrown when calling MillisecondsElapsed on a timer that hasn't stopped")
	}
}

func TestNanosecondsElapsed(t *testing.T) {
	timer := Timer{}
	timer.Start()
	// sleep for 1 second so that duration is at least 1,000,000,000 nanoseconds
	time.Sleep(time.Second)
	timer.Stop()

	duration, _ := timer.NanosecondsElapsed()
	if !(duration >= 1000000000 && duration < 2000000000) {
		t.Fatalf("expected duration to be greater than or equal to 1,000,000,000 but less than 2,000,000,000, instead it was %v", duration)
	}
	timer2 := Timer{}
	timer2.Start()
	_, err := timer2.NanosecondsElapsed()
	if err == nil {
		t.Fatalf("expected an error to be thrown when calling MillisecondsElapsed on a timer that hasn't stopped")
	}
}

func TestCompareTimesElapsed(t *testing.T) {
	timer := Timer{}
	timer.Start()
	time.Sleep(time.Second)
	timer.Stop()
	seconds, _ := timer.SecondsElapsed()
	milliseconds, _ := timer.MillisecondsElapsed()
	microseconds, _ := timer.MicrosecondsElapsed()
	nanoseconds, _ := timer.NanosecondsElapsed()

	if seconds*999 > milliseconds {
		t.Fatalf("Expected milliseconds to be ~1000 times larger than seconds. seconds: %v, milliseconds: %v", seconds, milliseconds)
	}
	if milliseconds*999 > microseconds {
		t.Fatalf("Expected microseconds to be ~1000 times larger than milliseconds.  milliseconds: %v, microseconds: %v", milliseconds, microseconds)
	}
	if microseconds*999 > nanoseconds {
		t.Fatalf("Expected nanoseconds to be ~1000 times larger than microseconds. microseconds: %v, nanoseconds: %v", microseconds, nanoseconds)
	}
}
