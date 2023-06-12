package list

type Stack[T interface{}] struct {
	list *LinkedList[T]
}

func NewStack[T interface{}]() *Stack[T] {
	return &Stack[T]{
		list: NewLinkedList[T](),
	}
}

func (s *Stack[T]) Push(item T) {
	s.list.InsertFirst(item)
}

func (s *Stack[T]) Pop() T {
	return s.list.DeleteFirst()
}

func (s *Stack[T]) IsEmpty() bool {
	return s.list.IsEmpty()
}
