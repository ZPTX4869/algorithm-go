package sorts

import "golang.org/x/exp/constraints"

func InsertSort[T constraints.Ordered](vals []T) {
	for i := 1; i < len(vals); i++ {
		insVal := vals[i]

		j := i 
		for ; j > 0 && insVal < vals[j-1]; j-- {
			vals[j] = vals[j-1]
		}
		vals[j] = insVal
	}
}
