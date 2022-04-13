package util

import "golang.org/x/exp/constraints"

func Sum[T constraints.Signed](vals []T) T {
	var sum T

	for i := 0; i < len(vals); i++ {
		sum += vals[i]
	}

	return sum
}