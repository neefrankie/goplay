package graph

import "golang.org/x/exp/constraints"

type Vertex[T constraints.Ordered] struct {
	data    T
	visited bool
}

func NewVertex[T constraints.Ordered](item T) *Vertex[T] {
	return &Vertex[T]{
		data:    item,
		visited: false,
	}
}
