package ch5

import "testing"

func Test_sum(t *testing.T) {
	t.Log(sum())
	t.Log(sum(3))
	t.Log(sum(1, 2, 3, 4))

	// To invoke a variadic function when the arguments are already in a slice,
	// place an ellipsis after the final argument.
	values := []int{1, 2, 3, 4}
	t.Log(sum(values...))
}
