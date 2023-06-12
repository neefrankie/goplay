package heap

import "sort"

func Sort(data sort.Interface) {
	size := data.Len() - 1

	for j := size/2 - 1; j >= 0; j-- {
		TrickleDown(data, j, size)
	}

	for j := size; j > 0; j-- {
		data.Swap(0, j)
		TrickleDown(data, 0, j-1)
	}
}
