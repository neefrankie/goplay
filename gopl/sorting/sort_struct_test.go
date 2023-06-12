package sorting

import (
	"sort"
	"testing"
)

func TestSortByAge(t *testing.T) {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	t.Logf("%v", people)

	sort.Sort(ByAge(people))

	t.Logf("%v", people)
}

func TestSortSlice(t *testing.T) {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	t.Logf("%v", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})

	t.Logf("%v", people)
}
