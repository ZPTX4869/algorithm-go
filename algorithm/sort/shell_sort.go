package sort

import "golang.org/x/exp/constraints"

func ShellSort[T constraints.Ordered](vals []T) {
	insertSort := func(vals []T, start, gap int) {
		for i := start + gap; i < len(vals); i += gap {
			insVal := vals[i]
			j := i
			for j >= start+gap && insVal < vals[j-gap] {
				vals[j] = vals[j-gap]
				j -= gap
			}
			vals[j] = insVal
		}
	}

	gap := len(vals) / 2
	for gap > 0 {
		for i := 0; i < len(vals)-gap; i++ {
			insertSort(vals, i, gap)
		}
		gap /= 2
	}
}
