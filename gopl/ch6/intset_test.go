package ch6

import (
	"fmt"
	"testing"
)

func TestIntSet(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())

	fmt.Println(x.Has(9), x.Has(123))
}
