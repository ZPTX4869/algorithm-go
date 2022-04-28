package heap

import "golang.org/x/exp/constraints"

type BinaryHeap[T any] struct {
	len  uint
	vals []T
	less func(T, T) bool
}

func New[T any](compare func(T, T) bool) BinaryHeap[T] {
	return BinaryHeap[T]{
		len: 0,
		// Leave T[0] unused for calculation simplicity
		vals: make([]T, 1),
		less: compare,
	}
}

func NewMin[T constraints.Ordered]() BinaryHeap[T] {
	return BinaryHeap[T]{
		len:  0,
		vals: make([]T, 1),
		less: func(v1, v2 T) bool {
			return v1 < v2
		},
	}
}

func NewMax[T constraints.Ordered]() BinaryHeap[T] {
	return BinaryHeap[T]{
		len:  0,
		vals: make([]T, 1),
		less: func(v1, v2 T) bool {
			return v1 > v2
		},
	}
}

func (h *BinaryHeap[T]) Len() uint {
	return h.len
}

func (h *BinaryHeap[T]) Empty() bool {
	return h.len == 0
}

func (h *BinaryHeap[T]) Push(val T) {
	h.len += 1
	h.vals = append(h.vals, val)

	// Heapify from bottom to top
	h.heapifyUp(h.len)
}

func (h *BinaryHeap[T]) Pop() T {
	if h.len == 0 {
		panic("the heap is empty")
	}

	h.vals[1], h.vals[h.len] = h.vals[h.len], h.vals[1]
	res := h.vals[h.len]
	h.vals = h.vals[:h.len]
	h.len -= 1

	h.heapifyDown(1)

	return res
}

func (h *BinaryHeap[T]) heapifyUp(idx uint) {
	for idx > 1 && h.less(h.vals[idx], h.vals[idx/2]) {
		h.vals[idx/2], h.vals[idx] = h.vals[idx], h.vals[idx/2]
		idx = idx / 2
	}
}

func (h *BinaryHeap[T]) heapifyDown(idx uint) {
	for 2*idx <= h.len {
		childIdx := 2 * idx
		if childIdx < h.len && h.less(h.vals[childIdx+1], h.vals[childIdx]) {
			childIdx += 1
		}

		if h.less(h.vals[idx], h.vals[childIdx]) {
			break
		}

		h.vals[idx], h.vals[childIdx] = h.vals[childIdx], h.vals[idx]
		idx = childIdx
	}
}
