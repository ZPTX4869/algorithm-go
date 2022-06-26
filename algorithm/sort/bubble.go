package sort

import "golang.org/x/exp/constraints"

func BubbleSort[T constraints.Ordered](vals []T) {
	for i := len(vals) - 1; i > 0; i-- {
		isSwap := false
		for j := 0; j < i; j++ {
			if vals[j] > vals[j+1] {
				vals[j], vals[j+1] = vals[j+1], vals[j]
				isSwap = true
			}
		}

		// If each next is larger than prev(already sorted), we can terminate loop.
		if isSwap == false {
			break
		}
	}
}
