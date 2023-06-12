package dynamic

// CashierChange makes n cents change with quarters, dimes, nickels,
// and pennies, and using het least total number of coins.
func CashierChange(coins []int, n int) []int {
	var counter = make([]int, len(coins))

	for i, c := range coins {
		for n >= c {
			counter[i] += 1
			n = n - c
		}
	}

	return counter
}
