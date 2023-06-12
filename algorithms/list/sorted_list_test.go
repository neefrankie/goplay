package list

import (
	"testing"
)

func TestNewSortedList(t *testing.T) {
	list := NewSortedList[int](func(current int, newItem int) bool {
		return current < newItem
	})

	list.Insert(10)
	list.Insert(20)
	list.Insert(30)
	list.Insert(5)

	list.Traverse(func(item int) {
		t.Log(item)
	})
}

func TestListInsertionSort(t *testing.T) {
	sorted := make([]int, 0)

	list := NewSortedList[int](func(current int, newItem int) bool {
		return current < newItem
	})

	list.Insert(10)
	list.Insert(20)
	list.Insert(30)
	list.Insert(5)

	for !list.IsEmpty() {
		item := list.Remove()
		sorted = append(sorted, item)
	}

	t.Logf("%v", sorted)
}
