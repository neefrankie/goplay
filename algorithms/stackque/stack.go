package stackque

import "errors"

type Stack[T interface{}] struct {
	arr     []T
	top     int
	maxSize int
}

func NewStack[T interface{}](size int) *Stack[T] {
	return &Stack[T]{
		arr:     make([]T, size),
		top:     -1,
		maxSize: size,
	}
}

// Push puts item on top of stack
func (s *Stack[T]) Push(x T) error {
	if s.IsFull() {
		return errors.New("stack is full")
	}
	s.top++
	s.arr[s.top] = x

	return nil
}

// Pop takes item from top of stack.
// Returned error is stack is empty.
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var ret T
		return ret, errors.New("underflow")
	}

	// Access item, decrement top.
	x := s.arr[s.top]
	s.top--
	return x, nil
}

// Peek inspects item at top.
func (s *Stack[T]) Peek() T {
	return s.arr[s.top]
}

func (s *Stack[T]) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack[T]) IsFull() bool {
	return s.top == s.maxSize-1
}

func (s *Stack[T]) Size() int {
	return s.top + 1
}

// Traverse the underline array from bottom to top.
func (s *Stack[T]) Traverse(visitor func(item T)) {
	for _, v := range s.arr {
		visitor(v)
	}
}
