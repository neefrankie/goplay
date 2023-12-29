package format

import (
	"testing"
	"time"
)

func TestAny(t *testing.T) {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	t.Log(Any(x))
	t.Log(Any(d))
	t.Log(Any([]int64{x}))
	t.Log(Any([]time.Duration{d}))
}
