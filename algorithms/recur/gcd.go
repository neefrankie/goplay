package recur

func EuclidGCD(m, n int) int {
	if n == 0 {
		return m
	}

	return EuclidGCD(n, m%n)
}
