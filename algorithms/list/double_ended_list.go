package list

// DoubleEndedList has a reference to the last link
// as well as to the first.
type DoubleEndedList[T interface{}] struct {
	head *Link[T]
	tail *Link[T]
}

func NewDoubleEndedList[T interface{}]() *DoubleEndedList[T] {
	return &DoubleEndedList[T]{
		head: nil,
		tail: nil,
	}
}

func (list *DoubleEndedList[T]) IsEmpty() bool {
	return list.head == nil
}

func (list *DoubleEndedList[T]) InsertFirst(item T) {
	link := NewLink(item)

	if list.IsEmpty() {
		list.insertInitial(link)
		return
	}

	link.next = list.head
	list.head = link
}

func (list *DoubleEndedList[T]) insertInitial(l *Link[T]) {
	list.head = l
	list.tail = l
}

func (list *DoubleEndedList[T]) InsertLast(item T) {
	link := NewLink(item)

	if list.IsEmpty() {
		list.insertInitial(link)
		return
	}

	list.tail.next = link
	list.tail = link
}

func (list *DoubleEndedList[T]) DeleteFirst() T {
	if list.IsEmpty() {
		var t T
		return t
	}

	item := list.head.item
	// If only one item.
	if list.head.next == nil {
		list.tail = nil
	}

	list.head = list.head.next
	return item
}

func (list *DoubleEndedList[T]) Traverse(visitor func(item T)) {
	if list.IsEmpty() {
		return
	}

	current := list.head
	for current != nil {
		visitor(current.item)
		current = current.next
	}
}
