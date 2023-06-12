package recur

import "testing"

func TestSumN(t *testing.T) {
	a := SumN(100)

	t.Logf("%d", a)
}

func TestFactorial(t *testing.T) {
	a := Factorial(10)

	t.Logf("%d", a)
}
