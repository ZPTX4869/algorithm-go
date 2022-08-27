package search

import "golang.org/x/exp/constraints"

func BinarySearch[T constraints.Ordered](vals []T, tar T) int {
	lp, rp := 0, len(vals)-1

	for lp <= rp {
		mid := lp + (rp-lp)/2

		if tar == vals[mid] {
			return mid
		} else if tar < vals[mid] {
			rp = mid - 1
		} else if tar > vals[mid] {
			lp = mid + 1
		}
	}

	return -1
}
