package list

import (
	"testing"
)

func TestNewDoubleEndedList(t *testing.T) {
	list := NewDoubleEndedList[int]()

	list.InsertFirst(20)
	list.InsertLast(30)
	list.InsertFirst(40)

	list.Traverse(func(item int) {
		t.Log(item)
	})

	list.DeleteFirst()

	list.Traverse(func(item int) {
		t.Log(item)
	})
}
