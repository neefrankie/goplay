package ch5

import "testing"

func Test_squares(t *testing.T) {
	f := squares()
	t.Logf("%d\n", f())
	t.Logf("%d\n", f())
	t.Logf("%d\n", f())
	t.Logf("%d\n", f())
}

func Test_topoSort(t *testing.T) {
	order := topoSort(prereqs)

	for i, v := range order {
		t.Logf("%d: %s\n", i, v)
	}
}
