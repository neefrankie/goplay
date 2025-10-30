package ch5

import "testing"

func Test_squares(t *testing.T) {
	f := squares()
	t.Logf("%d\n", f())
	t.Logf("%d\n", f())
	t.Logf("%d\n", f())
	t.Logf("%d\n", f())
}
