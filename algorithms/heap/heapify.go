package heap

import "sort"

func TrickleUp(data sort.Interface, k int) {
	for p := (k - 1) / 2; k > 0 && data.Less(p, k); p = (p - 1) / 2 {
		data.Swap(p, k)
		k = p
	}
}

func TrickleDown(data sort.Interface, k, n int) {
	for j := 2*k + 1; j <= n; j = 2*j + 1 {
		if j < n && data.Less(j, j+1) {
			j++
		}
		if data.Less(k, j) {
			data.Swap(k, j)
		}
		k = j
	}
}
