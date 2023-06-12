package hashing

import "errors"

type Hashable interface {
	HashKey() int
}

type IntHashable int

func (i IntHashable) HashKey() int {
	return int(i)
}

var NoItem = IntHashable(-1)

type HashTable struct {
	size int
	arr  []Hashable
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		size: size,
		arr:  make([]Hashable, size),
	}
}

func (h *HashTable) HashFunc(key int) int {
	return key % h.size
}

func (h *HashTable) LinearProbe(index int) int {
	index++
	index %= h.size // Wrap around if index exceeds array length.

	return index
}

func (h *HashTable) DoubleHash(index int) int {
	return 5 - index%5
}

func (h *HashTable) Insert(item Hashable) {
	key := item.HashKey()
	index := h.HashFunc(key)

	for h.arr[index] != nil && h.arr[index].HashKey() != -1 {
		index = h.LinearProbe(index)
	}

	h.arr[index] = item
}

func (h *HashTable) Delete(item Hashable) (Hashable, error) {
	hashKey := item.HashKey()
	index := h.HashFunc(hashKey)

	for h.arr[index] != nil {
		if h.arr[index].HashKey() == hashKey {
			temp := h.arr[index]
			h.arr[index] = NoItem
			return temp, nil
		}
		index = h.LinearProbe(index)
	}

	return nil, errors.New("not deleted")
}

func (h *HashTable) Find(key int) (Hashable, error) {
	index := h.HashFunc(key)

	for h.arr[index] != nil {
		if h.arr[index].HashKey() == key {
			return h.arr[index], nil
		}

		index = h.LinearProbe(index)
	}

	return nil, errors.New("not found")
}
