package pagee

import (
	"testing"
	"time"
)

func TestIntervalReadInt(t *testing.T) {
	enum := newEnum(1, 1, 3)
	slowEnum := intervalReadInt(1, enum)
	start := time.Now()

	for range slowEnum {
	}

	duration := time.Now().Sub(start) / time.Second
	seconds := int(duration)
	if seconds < 3 || seconds > 4 {
		t.Error("too long ? too short ", seconds)
	}

}
