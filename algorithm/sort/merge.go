package sort

import "golang.org/x/exp/constraints"

func MergeSort[T constraints.Ordered](vals []T) {
	sort(vals, 0, len(vals)-1)
}

func sort[T constraints.Ordered](vals []T, left, right int) {
	if left >= right {
		return
	}

	mid := (right-left)/2 + left

	sort(vals, left, mid)
	sort(vals, mid+1, right)
	merge(vals, left, mid, right)
}

func merge[T constraints.Ordered](vals []T, left, mid, right int) {
	temp := make([]T, 0)
	lp, rp := left, mid+1

	for lp <= mid && rp <= right {
		if vals[lp] <= vals[rp] {
			temp = append(temp, vals[lp])
			lp++
		} else {
			temp = append(temp, vals[rp])
			rp++
		}
	}

	if lp > mid {
		temp = append(temp, vals[rp:right+1]...)
	} else {
		temp = append(temp, vals[lp:mid+1]...)
	}

	copy(vals[left:right+1], temp)
}
