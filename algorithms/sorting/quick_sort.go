package sorting

import (
	"golang.org/x/exp/constraints"
)

func QuickSort[T constraints.Ordered](arr []T) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort[T constraints.Ordered](arr []T, p, r int) {
	if p < r {
		q := partition(arr, p, r)
		quickSort(arr, p, q-1)
		quickSort(arr, q+1, r)
	}
}
func partition[T constraints.Ordered](arr []T, p, r int) int {
	var x = arr[r]
	// Tracks the one larger than pivot.
	var i = p - 1
	// Loop until one less than r
	for j := p; j < r; j++ {
		// Find one smaller than x
		// [2, 4, 9, 8, 5, 7]
		if arr[j] <= x {
			// When j points to 9, i points to 4 and stops increase.
			// j continues to move till 5.
			// Then i increase by 1 to 9. Then 9 and 5 swapped.
			i += 1
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// When loops ends, the array will be [2, 4, 5, 8, 9, 7]. i is 2

	q := i + 1 // Partition point. Here i 3
	// Swap partition point with rightmost element. We will get [2, 3, 5, 7, 9, 8]
	arr[q], arr[r] = arr[r], arr[q]

	return q
}
