package array

import (
	"math"
	"sort"
	"testing"
)

func TestSortInts(t *testing.T) {
	s := []int{5, 2, 6, 3, 1, 4}

	sort.Ints(s)

	t.Logf("%v", s)
}

func TestSortFloats(t *testing.T) {
	s := []float64{5.2, -1.3, 0.7, -3.8, 2.6}

	sort.Float64s(s)
	t.Logf("%v", s)

	s = []float64{math.Inf(1), math.NaN(), math.Inf(-1), 0.0}
	sort.Float64s(s)
	t.Logf("%v", s)
}
