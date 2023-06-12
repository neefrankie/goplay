package recur

func Factorial(n int64) int64 {
	if n == 0 {
		return 1
	} else {
		return n * Factorial(n-1)
	}
}
