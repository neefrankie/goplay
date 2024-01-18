package ch5

// A variadic function is one that can be called with varying numbers of arguments.
// To declare a variadic function, the type of the final parameter is preceded by an
// ellipsis `...`, which indicates that the function may be called with any number
// of arguments of this type.
func sum(vals ...int) int {
	total := 0

	for _, val := range vals {
		total += val
	}

	return total
}

// Within the body of the function, the type of `vals` is an `[]int` slice.
// When sum is called, any number of values may be proviced for its `vals` parameter.
//
// Implicitly, the caller allocates an array, copies the arguments into it,
// and passes a slice of the entire array to the function.
//
// Althoug the `...int` parameter behaves like a slice within the function body,
// the type of a variadic function is distinct from the type of the function with
// an ordinary slice parameter.
