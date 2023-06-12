package tree

import "golang.org/x/exp/constraints"

type RBNode[T constraints.Ordered] struct {
	isRead bool
	item   T
	left   *RBNode[T]
	right  *RBNode[T]
}

func NewRBNode[T constraints.Ordered](item T, left, right *RBNode[T], isRead bool) *RBNode[T] {
	return &RBNode[T]{
		isRead: isRead,
		item:   item,
		left:   left,
		right:  right,
	}
}

// func rbInsert[T constraints.Ordered](node *RBNode[T], item T, fromLeft bool) *RBNode[T] {
// 	if node == nil {
// 		return NewRBNode(item, nil, nil, true)
// 	}

// }
