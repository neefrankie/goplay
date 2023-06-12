package ch2

import "testing"

func TestPopCount(t *testing.T) {
	var pc [256]byte

	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)

		t.Logf("pc[%d] + byte(%d&1) = %d", i/2, i, pc[i])
	}
}
