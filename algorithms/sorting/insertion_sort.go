package sorting

import (
	"golang.org/x/exp/constraints"
)

func InsertionSort[T constraints.Ordered](arr []T) []T {

	for i := 1; i < len(arr); i++ {

		var temp = arr[i]

		var j = i

		for j > 0 && arr[j-1] >= temp {

			arr[j] = arr[j-1]
			j--
		}
		arr[j] = temp
	}

	return arr
}
