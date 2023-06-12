package search

func BinarySearch(x int, arr []int) int {
	var i = 0
	var j = len(arr) - 1

	for i < j {
		m := (i + j) / 2

		if x == arr[m] {
			return m
		} else if x < arr[m] {
			j = m
		} else {
			i = m + 1
		}
	}

	return -1
}
