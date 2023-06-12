package graph

const maxVerts = 20
const infinity = 1000000

type WeightedGraph struct {
	vertexList []*Vertex[string]
	adjMat     [][]int
	nVerts     int
}

func NewWeightedGraph() *WeightedGraph {
	adjMat := make([][]int, maxVerts)
	for i := 0; i < maxVerts; i++ {
		adjMat[i] = make([]int, maxVerts)
		for k := 0; k < maxVerts; k++ {
			adjMat[i][k] = infinity
		}
	}

	return &WeightedGraph{
		vertexList: make([]*Vertex[string], maxVerts),
		adjMat:     adjMat,
		nVerts:     0,
	}
}

func (g *WeightedGraph) AddVertex(label string) {
	g.vertexList[g.nVerts] = NewVertex(label)
	g.nVerts++
}

func (g *WeightedGraph) AddEdge(start, end, weight int) {
	g.adjMat[start][end] = weight
	g.adjMat[end][start] = weight
}

func (g *WeightedGraph) WeightedMST(visitor func(src, dest string)) {
	currentVert := 0
	nTree := 0
	queue := newPriorityQueue()

	// While not all vertices in tree, put currentVert in tree.
	for nTree < g.nVerts-1 {
		g.vertexList[currentVert].visited = true
		nTree++

		// Insert edges adjacent to currentVert into priority queue.
		for j := 0; j < g.nVerts; j++ {
			if j == currentVert {
				continue
			}

			if g.vertexList[j].visited {
				continue
			}

			distance := g.adjMat[currentVert][j]
			if distance == infinity {
				continue
			}

			queIdx := queue.findDex(j)
			if queIdx != -1 {
				tempEdge := queue.peekN(queIdx)
				oldDist := tempEdge.distance
				if oldDist > distance {
					queue.removeN(queIdx)
					queue.insert(weightedEdge{
						srcVert:  currentVert,
						destVert: j,
						distance: distance,
					})
				}
			} else {
				queue.insert(weightedEdge{
					srcVert:  currentVert,
					destVert: j,
					distance: distance,
				})
			}
		}

		if queue.count() == 0 {
			panic("graph not connected")
		}

		edge := queue.removeMin()
		srcIdx := edge.srcVert
		currentVert = edge.destVert

		visitor(g.vertexList[srcIdx].data, g.vertexList[currentVert].data)
	}

	for j := 0; j < g.nVerts; j++ {
		g.vertexList[j].visited = false
	}
}
