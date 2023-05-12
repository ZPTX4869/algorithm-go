package sort

import "golang.org/x/exp/constraints"

func ShellSort[T constraints.Ordered](vals []T) {
	insertSort := func(gap int) {
		for i := gap; i < len(vals); i += 1 {
			temp := vals[i]
			j := i
			for ; j >= gap && vals[j-gap] >= temp; j -= gap {
				vals[j] = vals[j-gap]
			}
			vals[j] = temp
		}
	}

	for gap := len(vals); gap > 0; gap /= 2 {
		insertSort(gap)
	}
}
