package pagee

import (
	"testing"
	"time"
)

func testDuration(t *testing.T, min time.Duration, max time.Duration, fn func()) {
	start := time.Now()
	fn()
	duration := time.Now().Sub(start) / time.Second
	if duration < min || duration > max {
		t.Error("too long ? too short ", duration)
	}
}
