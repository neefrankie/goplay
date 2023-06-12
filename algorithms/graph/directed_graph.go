package graph

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"log"
)

// DirectedGraph has edges with direction.
// In a directed graph you can proceed only one way along an edge.
// It differs from a non-directed graph in that
// an edge in a directed graph has only one entry in the adjacency matrix.
// For a weighted graph, every cell in the adjacency matrix conveys
// unique information. The halves aare not mirror images.
type DirectedGraph[T constraints.Ordered] struct {
	max        int
	vertexList []*Vertex[T]
	adjMat     [][]int
	count      int
}

func NewDirectedGraph[T constraints.Ordered]() *DirectedGraph[T] {
	max := 20

	adjMat := make([][]int, max)
	for i := 0; i < max; i++ {
		adjMat[i] = make([]int, max)
	}

	return &DirectedGraph[T]{
		max:        max,
		vertexList: make([]*Vertex[T], 20),
		adjMat:     adjMat,
		count:      0,
	}
}

func (d *DirectedGraph[T]) AddVertex(item T) {
	d.vertexList[d.count] = NewVertex(item)
	d.count++
}

func (d *DirectedGraph[T]) AddEdge(start, end int) {
	d.adjMat[start][end] = 1
}

func (d *DirectedGraph[T]) TopologicalSort() []T {
	var sorted = make([]T, 0)

	d.display()
	log.Println()

	for d.count > 0 {
		curIdx := d.noSuccessorsVertex()
		if curIdx == -1 {
			panic("Graph has cycles")
		}

		sorted = append(sorted, d.vertexList[curIdx].data)

		log.Printf("Will delete vertex %v", d.vertexList[curIdx].data)
		d.deleteVertex(curIdx)
		log.Printf("After vertex %v deleted:\n", d.vertexList[curIdx].data)
		d.display()
	}

	return sorted
}

// noSuccessorsVertex returns the index of vertex without successors,
// or -1 if no such vertex.
func (d *DirectedGraph[T]) noSuccessorsVertex() int {
	var hasEdge bool

	// Loop over all vertices.
	for row := 0; row < d.count; row++ {
		hasEdge = false
		// Check if a vertex is linked to any other vertex.
		for col := 0; col < d.count; col++ {
			if d.adjMat[row][col] > 0 {
				hasEdge = true
				break
			}
		}
		// If we find a vertex without link, stop.
		if !hasEdge {
			return row
		}
	}

	// No vertex without link.
	return -1
}

// deleteVertex from the graph. When index is not the last element,
// you have to move following rows up, and columns left:
//     A    B    C    D
// ---------------------
// A | 0    1    0    0
// B | 0    0    1    1
// C | 0    0    0    0
// D | 0    0    0    0
//
// For the graph:
// A -> B -> C
//     |
//     V
//     D
// When deleting C, you have to move row D up, and column D left.
func (d *DirectedGraph[T]) deleteVertex(delIdx int) {
	if delIdx != d.count-1 {
		for i := delIdx; i < d.count-1; i++ {
			d.vertexList[i] = d.vertexList[i+1]
		}

		for row := delIdx; row < d.count-1; row++ {
			d.moveRowUp(row, d.count)
		}

		for col := delIdx; col < d.count-1; col++ {
			d.moveColLeft(col, d.count-1) // The rows are reduced by 1.
		}
	}
	d.count--
}

func (d *DirectedGraph[T]) display() {
	for i := 0; i < d.count; i++ {
		fmt.Printf("%v: %v\n", d.vertexList[i].data, d.adjMat[i][:d.count])
	}
}

func (d *DirectedGraph[T]) moveRowUp(row int, colLen int) {
	for col := 0; col < colLen; col++ {
		d.adjMat[row][col] = d.adjMat[row+1][col]
	}
}

func (d *DirectedGraph[T]) moveColLeft(col int, rowLen int) {
	for row := 0; row < rowLen; row++ {
		d.adjMat[row][col] = d.adjMat[row][col+1]
	}
}
