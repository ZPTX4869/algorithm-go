package sorts

import "golang.org/x/exp/constraints"

func HeapSort[T constraints.Ordered](vals []T) {
	if len(vals) == 1 {
		return
	}
}
