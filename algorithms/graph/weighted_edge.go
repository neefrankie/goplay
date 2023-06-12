package graph

type weightedEdge struct {
	srcVert  int
	destVert int
	distance int
}

type priorityQueue struct {
	max  int
	arr  []weightedEdge // array in sorted order, from max at 0 to min at size-1
	size int
}

func newPriorityQueue() *priorityQueue {
	max := 20
	return &priorityQueue{
		max:  max,
		arr:  make([]weightedEdge, max),
		size: 0,
	}
}

func (q *priorityQueue) insert(item weightedEdge) {
	j := 0

	for ; j < q.size; j++ {
		if item.distance >= q.arr[j].distance {
			break
		}
	}

	for k := q.size; k >= j; k-- {
		q.arr[k+1] = q.arr[k]
	}

	q.arr[j] = item
}

func (q *priorityQueue) removeMin() weightedEdge {
	q.size--
	return q.arr[q.size]
}

func (q *priorityQueue) removeN(n int) {
	for j := n; j < q.size; j++ {
		q.arr[j] = q.arr[j+1]
	}
	q.size--
}

func (q *priorityQueue) peekMin() weightedEdge {
	return q.arr[q.size-1]
}

func (q *priorityQueue) peekN(n int) weightedEdge {
	return q.arr[n]
}

func (q *priorityQueue) count() int {
	return q.size
}

func (q *priorityQueue) isEmpty() bool {
	return q.size == 0
}

func (q *priorityQueue) findDex(destVert int) int {
	for j := 0; j < q.size; j++ {
		if q.arr[j].destVert == destVert {
			return j
		}
	}

	return -1
}
