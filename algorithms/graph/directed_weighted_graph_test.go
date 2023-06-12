package graph

import (
	"testing"
)

func TestNewDirectedWeightedGraph(t *testing.T) {
	g := NewDirectedWeightedGraph()

	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("D")
	g.AddVertex("E")

	g.AddEdge(0, 1, 50)
	g.AddEdge(0, 3, 80)
	g.AddEdge(1, 2, 60)
	g.AddEdge(1, 3, 90)
	g.AddEdge(2, 4, 40)
	g.AddEdge(3, 2, 20)
	g.AddEdge(3, 4, 70)
	g.AddEdge(4, 1, 50)

	sPath := g.Path()

	t.Logf("%v", sPath)
}
