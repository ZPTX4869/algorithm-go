package sort

import "golang.org/x/exp/constraints"

func HeapSort[T constraints.Ordered](vals []T) {
	n := len(vals)
	if n < 2 {
		return
	}

	for i := n; i >= 0; i-- {
		heapify(vals, (n-1)/2, n)
	}

	for i := n - 1; i >= 0; i-- {
		vals[0], vals[i] = vals[i], vals[0]
		heapify(vals, 0, i)
	}
}

func heapify[T constraints.Ordered](vals []T, root, n int) {
	idxMax := root
	lchild := root*2 + 1
	rchild := root*2 + 2

	if lchild < n && vals[lchild] > vals[idxMax] {
		idxMax = lchild
	}

	if rchild < n && vals[rchild] > vals[idxMax] {
		idxMax = rchild
	}

	if idxMax != root {
		vals[root], vals[idxMax] = vals[idxMax], vals[root]
		heapify(vals, idxMax, n)
	}
}
