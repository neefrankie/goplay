package graph

import (
	"fmt"
	"strconv"
)

type distanceParent struct {
	distance      int
	parentVertIdx int
}

type PathItem struct {
	Vertex       *Vertex[string]
	Distance     int
	ParentVertex *Vertex[string]
}

func (i PathItem) String() string {
	var dis string
	if i.Distance == infinity {
		dis = "inf"
	} else {
		dis = strconv.FormatInt(int64(i.Distance), 10)
	}

	return fmt.Sprintf("%s=%s(%s)", i.Vertex.data, dis, i.ParentVertex.data)
}

type DirectedWeightedGraph struct {
	vertexList []*Vertex[string]
	adjMat     [][]int
	nVerts     int
	nTree      int
}

func NewDirectedWeightedGraph() *DirectedWeightedGraph {
	adjMat := make([][]int, maxVerts)
	for i := 0; i < maxVerts; i++ {
		adjMat[i] = make([]int, maxVerts)
		for k := 0; k < maxVerts; k++ {
			adjMat[i][k] = infinity
		}
	}

	return &DirectedWeightedGraph{
		vertexList: make([]*Vertex[string], maxVerts),
		adjMat:     adjMat,
		nVerts:     0,
		nTree:      0,
	}
}

func (g *DirectedWeightedGraph) AddVertex(label string) {
	g.vertexList[g.nVerts] = NewVertex(label)
	g.nVerts++
}

func (g *DirectedWeightedGraph) AddEdge(start, end, weight int) {
	g.adjMat[start][end] = weight
}

func (g *DirectedWeightedGraph) Path() []PathItem {
	startTree := 0
	g.vertexList[startTree].visited = true
	g.nTree = 1

	shortestPath := make([]distanceParent, g.nVerts)

	// Copy the first row of distance from adjacency matrix to shorted path.
	for j := 0; j < g.nVerts; j++ {
		tempDist := g.adjMat[startTree][j]
		shortestPath[j] = distanceParent{
			distance:      tempDist,
			parentVertIdx: startTree,
		}
	}

	for g.nTree < g.nVerts {
		idxMin := g.getMinIdx(shortestPath)
		minDist := shortestPath[idxMin].distance

		if minDist == infinity {
			panic("there are unreachable vertices")
		}

		g.vertexList[idxMin].visited = true
		g.nTree++
		g.adjustShortestPath(shortestPath, idxMin)
	}

	items := make([]PathItem, 0)
	for j := 0; j < g.nVerts; j++ {
		item := PathItem{
			Vertex:       g.vertexList[j],
			Distance:     shortestPath[j].distance,
			ParentVertex: g.vertexList[shortestPath[j].parentVertIdx],
		}

		items = append(items, item)
	}

	return items
}

func (g *DirectedWeightedGraph) getMinIdx(sPath []distanceParent) int {
	minDist := infinity
	idxMin := 0

	for j := 1; j < g.nVerts; j++ {
		if !g.vertexList[j].visited && sPath[j].distance < minDist {
			minDist = sPath[j].distance
			idxMin = j
		}
	}

	return idxMin
}

func (g *DirectedWeightedGraph) adjustShortestPath(sPath []distanceParent, row int) {
	startToCurrent := sPath[row].distance
	column := 1

	for column < g.nVerts {
		if g.vertexList[column].visited {
			column++
			continue
		}

		currentToFringe := g.adjMat[row][column]
		// Newly calculated distance from start vertex to a fringe vertex.
		startToFringe := startToCurrent + currentToFringe
		// Existing distance from start vertex to this fringe vertex.
		sPathDist := sPath[column].distance

		if startToFringe < sPathDist {
			sPath[column].parentVertIdx = row
			sPath[column].distance = startToFringe
		}
		column++
	}
}
