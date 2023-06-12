package array

func CountNumbers1(n int) int {

	count := 0
	mul := 10

	for n >= 0 {
		count++
		n = n - mul
		mul *= 10
	}

	return count
}

func CountNumbers2(n int) int {
	count := 0
	for ; n > 0; n /= 10 {
		count++
	}

	return count
}

func CountDigits(n int) int {
	if n < 10 {
		return 1
	}

	return CountDigits(n/10) + 1
}
