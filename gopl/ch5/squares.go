package ch5

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
