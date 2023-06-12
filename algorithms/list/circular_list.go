package list

import "golang.org/x/exp/constraints"

type CircularList[T constraints.Ordered] struct {
	head *Link[T]
	tail *Link[T]
}

func NewCircularList[T constraints.Ordered]() *CircularList[T] {
	return &CircularList[T]{
		head: nil, // Point to the last inserted element.
		tail: nil, // Point to the initially inserted element
	}
}

func (list *CircularList[T]) IsEmpty() bool {
	return list.head == nil
}

func (list *CircularList[T]) Insert(item T) {
	t := NewLink(item)

	if list.IsEmpty() {
		t.next = t
		list.head = t
		list.tail = t
		return
	}

	list.head.next = t
	list.head = t

	t.next = list.tail
}

func (list *CircularList[T]) GetHead() *Link[T] {
	return list.head
}

func (list *CircularList[T]) Traverse(visitor func(item T)) {
	if list.IsEmpty() {
		return
	}

	current := list.head.next
	for current != list.head {
		visitor(current.item)
		current = current.next
	}

	visitor(list.head.item)
}

func JosephusProblem(size int, every int) int {
	list := NewCircularList[int]()

	for i := 1; i <= size; i++ {
		list.Insert(i)
	}

	current := list.GetHead()
	for current != current.next {
		for i := 1; i < every; i++ {
			current = current.next
		}

		current.next = current.next.next
	}

	return current.item
}
