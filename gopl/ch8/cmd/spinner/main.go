package main

import (
	"fmt"
	"gopl/ch8/spinner"
	"time"
)

func main() {
	go spinner.Spinner(100 * time.Millisecond)
	const n = 45
	fibN := spinner.Fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
