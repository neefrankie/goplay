package graph

import (
	"goalgorithms/stackque"
	"golang.org/x/exp/constraints"
)

type Graph[T constraints.Ordered] struct {
	max        int
	vertexList []*Vertex[T]
	adjMat     [][]int
	count      int
}

func NewGraph[T constraints.Ordered]() *Graph[T] {
	max := 20

	adjMat := make([][]int, max)
	for i := 0; i < max; i++ {
		adjMat[i] = make([]int, max)
	}

	return &Graph[T]{
		max:        max,
		vertexList: make([]*Vertex[T], 20),
		adjMat:     adjMat,
		count:      0,
	}
}

func (g *Graph[T]) AddVertex(item T) {
	g.vertexList[g.count] = NewVertex(item)
	g.count++
}

func (g *Graph[T]) AddEdge(start, end int) {
	g.adjMat[start][end] = 1
	g.adjMat[end][start] = 1
}

func (g *Graph[T]) DepthFirstTraverse(visitor func(item T)) {
	stack := stackque.NewStack[int](g.max)

	g.vertexList[0].visited = true
	visitor(g.vertexList[0].data)
	_ = stack.Push(0)

	for !stack.IsEmpty() {
		v := g.getAdjUnvisitedVertex(stack.Peek())
		if v == -1 {
			_, _ = stack.Pop()
		} else {
			g.vertexList[v].visited = true
			visitor(g.vertexList[v].data)
			_ = stack.Push(v)
		}
	}

	for i := 0; i < g.count; i++ {
		g.vertexList[i].visited = false
	}
}

func (g *Graph[T]) BreadthFirstTraverse(visitor func(item T)) {
	q := stackque.NewQueue[int](g.max)

	g.vertexList[0].visited = true
	visitor(g.vertexList[0].data)
	_ = q.Enqueue(0)

	for !q.IsEmpty() {
		curVertIdx, _ := q.Dequeue()
		adjVertIdx := g.getAdjUnvisitedVertex(curVertIdx)

		for adjVertIdx != -1 {
			g.vertexList[adjVertIdx].visited = true
			visitor(g.vertexList[adjVertIdx].data)
			_ = q.Enqueue(adjVertIdx)

			adjVertIdx = g.getAdjUnvisitedVertex(curVertIdx)
		}
	}

	for i := 0; i < g.count; i++ {
		g.vertexList[i].visited = false
	}
}

func (g *Graph[T]) MST(visitor func(current T, next T)) {
	stack := stackque.NewStack[int](g.max)

	g.vertexList[0].visited = true
	_ = stack.Push(0)

	for !stack.IsEmpty() {
		curIdx := stack.Peek()
		adjIdx := g.getAdjUnvisitedVertex(curIdx)
		if adjIdx == -1 {
			_, _ = stack.Pop()
		} else {
			g.vertexList[adjIdx].visited = true
			_ = stack.Push(adjIdx)
			visitor(g.vertexList[curIdx].data, g.vertexList[adjIdx].data)
		}
	}

	for i := 0; i < g.count; i++ {
		g.vertexList[0].visited = false
	}
}

func (g *Graph[T]) getAdjUnvisitedVertex(v int) int {
	for i := 0; i < g.count; i++ {
		if g.adjMat[v][i] == 1 && !g.vertexList[i].visited {
			return i
		}
	}

	return -1
}
