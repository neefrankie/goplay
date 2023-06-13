package list

import "testing"

func TestMyLinkedList(t *testing.T) {
	myLinkedList := Constructor()
	myLinkedList.AddAtHead(1)

	myLinkedList.print()

	myLinkedList.AddAtTail(3)

	myLinkedList.print()

	myLinkedList.AddAtIndex(1, 2)

	myLinkedList.print()

	t.Logf("%d", myLinkedList.Get(1))

	myLinkedList.DeleteAtIndex(1)
	myLinkedList.print()

	t.Logf("%d", myLinkedList.Get(1))

	myLinkedList.DeleteAtIndex(0)
	myLinkedList.print()
}
