package sort

import "golang.org/x/exp/constraints"

func SelectSort[T constraints.Ordered](vals []T) {
	for i := 0; i < len(vals); i++ {
		idxMin := i

		for j := i+1; j < len(vals); j++ {
			if vals[j] < vals[idxMin] {
				idxMin = j
			}
		}
		
		vals[i], vals[idxMin] = vals[idxMin], vals[i]
	}
}