package graph

import (
	"testing"
)

func TestDirectedGraph_TopologicalSort(t *testing.T) {
	g := NewDirectedGraph[string]()
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("D")
	g.AddVertex("E")
	g.AddVertex("F")
	g.AddVertex("G")
	g.AddVertex("H")

	g.AddEdge(0, 3)
	g.AddEdge(0, 4)
	g.AddEdge(1, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)
	g.AddEdge(4, 6)
	g.AddEdge(5, 7)
	g.AddEdge(6, 7)

	sorted := g.TopologicalSort()

	t.Logf("%v", sorted)
}
