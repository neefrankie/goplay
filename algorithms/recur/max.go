package recur

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](items []T, lo, hi int) T {
	if lo == hi {
		return items[lo]
	}

	m := (lo + hi) / 2

	leftMax := Max(items, lo, m)
	rightMax := Max(items, m+1, hi)

	if leftMax > rightMax {
		return leftMax
	} else {
		return rightMax
	}
}
