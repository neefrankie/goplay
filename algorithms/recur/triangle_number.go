package recur

// SumN is the sum of first n integers.
func SumN(n int) int {
	if n == 1 {
		return 1
	} else {
		return n + SumN(n-1)
	}
}
