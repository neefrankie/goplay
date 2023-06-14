package list

import (
	"testing"
)

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

func Test_removeNthFromEnd(t *testing.T) {
	one := ListNode{Val: 1}
	two := ListNode{Val: 2}
	three := ListNode{Val: 3}
	four := ListNode{Val: 4}
	five := ListNode{Val: 5}

	one.Next = &two
	two.Next = &three
	three.Next = &four
	four.Next = &five

	head := removeNthFromEnd(&one, 2)

	head.traverse(func(v int) {
		t.Logf("%d ->", v)
	})

	one = ListNode{Val: 1}

	head = removeNthFromEnd(&one, 1)
	head.traverse(func(v int) {
		t.Logf("\n%d ->", v)
	})

	two = ListNode{Val: 2}
	one.Next = &two

	head = removeNthFromEnd(head, 1)
	head.traverse(func(v int) {
		t.Logf("\n%d ->", v)
	})
}
