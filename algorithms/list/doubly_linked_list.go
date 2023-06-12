package list

type DoublyLinkedList[T interface{}] struct {
	head *DoublyLink[T]
	tail *DoublyLink[T]
}

func NewDoublyLinkedList[T interface{}]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		head: nil,
		tail: nil,
	}
}

func (list *DoublyLinkedList[T]) IsEmpty() bool {
	return list.head == nil
}

func (list *DoublyLinkedList[T]) InsertFirst(item T) {
	link := NewDoublyLink(item)

	if list.IsEmpty() {
		list.tail = link
	} else {
		list.head.previous = link
	}

	link.next = list.head
	list.head = link
}

func (list *DoublyLinkedList[T]) InsertLast(item T) {
	link := NewDoublyLink(item)

	// If the list is empty, head should point to the new link.
	if list.IsEmpty() {
		list.head = link
	} else {
		// Tail's next point to this new link.
		list.tail.next = link
		// The new link point back to current tail.
		link.previous = list.tail
	}

	// Move tail to point to the new link.
	list.tail = link
}

func (list *DoublyLinkedList[T]) DeleteFirst() T {
	// Empty
	if list.head == nil {
		var t T
		return t
	}

	t := list.head
	// If there's only 1 element, the last one should be set to null
	if list.head.next == nil {
		list.tail = nil
	} else {
		list.head.next.previous = nil
	}

	list.head = list.head.next

	return t.item
}

func (list *DoublyLinkedList[T]) DeleteLast() T {
	// If there's no element.
	if list.tail == nil {
		var t T
		return t
	}

	t := list.tail
	if list.head.next == nil {
		list.head = nil
	} else {
		list.tail.previous.next = nil
	}

	list.tail = list.tail.previous

	return t.item
}
