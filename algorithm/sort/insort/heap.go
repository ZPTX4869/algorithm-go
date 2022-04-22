package insort

import "golang.org/x/exp/constraints"

func HeapSort[T constraints.Ordered](vals []T) {
	n := len(vals)
	if n < 2 {
		return
	}

	for i := (n - 1) / 2; i >= 0; i-- {
		// heapifyUp(vals, i)
		heapifyDown(vals, i, n)
	}

	for i := n - 1; i >= 0; i-- {
		vals[0], vals[i] = vals[i], vals[0]
		heapifyDown(vals, 0, i)
	}
}

// func heapifyUp[T constraints.Ordered](vals []T, idx int) {
// 	for idx > 0 && vals[idx] > vals[(idx-1)/2] {
// 		vals[idx], vals[(idx-1)/2] = vals[(idx-1)/2], vals[idx]
// 		idx = (idx - 1) / 2
// 	}
// }

func heapifyDown[T constraints.Ordered](vals []T, parent, end int) {
	for parent*2+1 < end {
		child := parent*2 + 1
		if child+1 < end && vals[child+1] > vals[child] {
			child += 1
		}
		if vals[parent] > vals[child] {
			break
		}
		vals[parent], vals[child] = vals[child], vals[parent]
		parent = child
	}
}
