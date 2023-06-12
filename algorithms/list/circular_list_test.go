package list

import (
	"testing"
)

func TestCircularList_Traverse(t *testing.T) {
	l := NewCircularList[int]()

	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Insert(4)
	l.Insert(5)
	l.Insert(6)
	l.Insert(7)
	l.Insert(8)
	l.Insert(9)

	l.Traverse(func(item int) {
		t.Logf("%d ", item)
	})
}

func TestJosephusProblem(t *testing.T) {
	survivor := JosephusProblem(9, 5)

	t.Logf("%d", survivor)
}
