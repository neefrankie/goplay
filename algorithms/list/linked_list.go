package list

type LinkedList[T interface{}] struct {
	head *Link[T]
}

func NewLinkedList[T interface{}]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (list *LinkedList[T]) IsEmpty() bool {
	return list.head == nil
}

func (list *LinkedList[T]) InsertFirst(item T) {
	link := NewLink(item)
	link.next = list.head
	list.head = link
}

func (list *LinkedList[T]) DeleteFirst() T {
	t := list.head
	list.head = list.head.next

	if t == nil {
		var item T
		return item
	}

	return t.item
}

func (list *LinkedList[T]) Find(by func(item T) bool) T {

	for current := list.head; current != nil; current = current.next {
		if by(current.item) {
			return current.item
		}

		current = current.next
	}

	var item T
	return item
}

func (list *LinkedList[T]) Delete(by func(item T) bool) T {

	var current = list.head
	var previous = list.head

	for current != nil {
		if by(current.item) {
			break
		}

		previous = current
		current = current.next
	}

	// If reached the end of list, or the list is empty
	if current == nil {
		var item T
		return item
	}

	if current == list.head {
		list.head = list.head.next
	} else {
		previous.next = current.next
	}

	return current.item
}

func (list *LinkedList[T]) Traverse(visitor func(item T)) {
	current := list.head

	for current != nil {
		visitor(current.item)
		current = current.next
	}
}

func (list *LinkedList[T]) GetFirst() *Link[T] {
	return list.head
}

func (list *LinkedList[T]) SetFirst(l *Link[T]) {
	list.head = l
}

func (list *LinkedList[T]) Iterator() *Iterator[T] {
	return NewIterator[T](list)
}

func (list *LinkedList[T]) Reverse() {
	list.head = Reverse(list.head)
}
