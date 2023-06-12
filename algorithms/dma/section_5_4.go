package dma

import "log"

// RecurMultiply is a recursive algorithm for computing nx whenever
// n is a positive integer and x is an integer, using just addition.
// Answer to exercise 7 on page 391.
func RecurMultiply(n, x int64) int64 {
	if n <= 0 {
		log.Fatalf("n must be a positive integer, got %d", n)
	}

	if n == 1 {
		return x
	}

	return RecurMultiply(n-1, x) + x
}

// RecurSumN is a recursive algorithm form finding the sum of the
// first n positive integer.
// Answer to exercise 8 on page 391
func RecurSumN(n int64) int64 {
	if n <= 0 {
		log.Fatalf("n must be a positive integer, got %d", n)
	}

	if n == 1 {
		return n
	}

	return n + RecurSumN(n-1)
}

// RecurSumOddN is a recursive algorithm for finding the sum of the first
// n odd positive integers.
// Answer to exercise 9 on page 391.
func RecurSumOddN(n int64) int64 {
	if n <= 0 {
		log.Fatalf("n must be a positive integer, got %d", n)
	}

	if n == 1 {
		return n
	}

	return (2*n - 1) + RecurSumOddN(n-1)
}

// RecurMax is a recursive algorithm for finding the maximum of a
// finite set of integers.
// The maximum of n integers is the larger of
// * the last integer in the list
// * and the maximum of the first n-1 integers in the list.
// Answer to exercise 10 on page 391.
func RecurMax(arr []int) int {
	n := len(arr)
	if n == 1 {
		return arr[0]
	}

	last := arr[n-1]
	headMax := RecurMax(arr[0 : n-1])
	if last > headMax {
		return last
	} else {
		return headMax
	}
}
