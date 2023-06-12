package graph

import (
	"testing"
)

func TestNewGraph(t *testing.T) {
	g := NewGraph[string]()

	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("D")
	g.AddVertex("E")

	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(0, 3)
	g.AddEdge(3, 4)

	t.Log("Depth first traversal")
	g.DepthFirstTraverse(func(item string) {
		t.Logf("%s", item)
	})

	t.Log("Breadth first traversal")
	g.BreadthFirstTraverse(func(item string) {
		t.Logf("%s", item)
	})
}

func TestMST(t *testing.T) {
	g := NewGraph[string]()

	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("D")
	g.AddVertex("E")

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 3)
	g.AddEdge(0, 4)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)

	g.MST(func(current string, next string) {
		t.Logf("%s-%s", current, next)
	})
}
