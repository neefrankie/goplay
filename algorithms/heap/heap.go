package heap

import "sort"

type Heap struct {
	heapArray []int
	size      int
}

func New() *Heap {
	return &Heap{
		heapArray: []int{},
		size:      0,
	}
}

func (h *Heap) Insert(weight int) {
	h.heapArray = append(h.heapArray, weight)
	TrickleUp(sort.IntSlice(h.heapArray), h.size)
	h.size++
}

func (h *Heap) Remove() int {
	root := h.heapArray[0]
	h.size--
	h.heapArray[0] = h.heapArray[h.size]
	TrickleDown(sort.IntSlice(h.heapArray), 0, h.size-1)

	return root
}
