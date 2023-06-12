package elementary

import (
	"fmt"
	"math/rand"
)

func SieveOfEratosthenes(n int) {
	var a = make([]int, n)
	for i := 2; i < n; i++ {
		a[i] = 1
	}

	for i := 2; i < n; i++ {
		if a[i] == 1 {
			for j := i; i*j < n; j++ {
				a[i*j] = 0
			}
		}
	}

	for i := 2; i < n; i++ {
		if a[i] == 1 {
			fmt.Printf("%4d", i)
		}
	}
}

// CoinFlipping show the histogram of coin flipping frequency.
func CoinFlipping(count int, round int) {

	var freq = make([]int, count)

	var heads = func() bool {
		return rand.Intn(2) > 0
	}

	for i := 0; i < round; i++ {
		cnt := 0
		for j := 0; j < count; j++ {
			if heads() {
				cnt++
			}
		}
		freq[cnt]++
	}

	for j := 0; j < count; j++ {
		fmt.Printf("%2d", j)
		for i := 0; i < freq[j]; i += 10 {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
