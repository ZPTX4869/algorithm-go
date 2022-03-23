package sorts

import "golang.org/x/exp/constraints"

func ShellSort[T constraints.Ordered](vals []T) {
	n := len(vals)

	for interval := n/2; interval > 0; interval /= 2 {
		for i := interval; i < n; i++ {
			insVal := vals[i]

			j := i
			for ; j > interval - 1 && insVal < vals[j-interval]; j -= interval {
				vals[j] = vals[j-interval]
			}
			vals[j] = insVal
		}
	} 
}