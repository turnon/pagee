package pagee

import (
	"testing"
	"time"
)

func TestIntervalReadInt(t *testing.T) {
	enum := newEnum(1, 1, 3)
	slowEnum := intervalReadInt(1, enum)

	testDuration(t, 3, 4, func() {
		for range slowEnum {
		}
	})
}

func testDuration(t *testing.T, min time.Duration, max time.Duration, fn func()) {
	start := time.Now()
	fn()
	duration := time.Now().Sub(start) / time.Second
	if duration < min || duration > max {
		t.Error("too long ? too short ", duration)
	}
}
