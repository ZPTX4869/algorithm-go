package sort

import "golang.org/x/exp/constraints"

func MergeSort[T constraints.Ordered](vals []T) {
	if len(vals) == 1 {
		return
	}
	
	sort(0, len(vals)-1, vals)
}

func sort[T constraints.Ordered](left, right int, vals []T) {
	if left >= right {
		return
	}

	mid := (right-left) / 2 + left

	sort(left, mid, vals)
	sort(mid+1, right, vals)
	merge(left, mid, right, vals)
}

func merge[T constraints.Ordered](left, mid, right int, vals []T) {
	res := make([]T, 0)

	lp, rp := left, mid+1
	for lp <= mid && rp <= right {
		if vals[lp] < vals[rp] {
			res = append(res, vals[lp])
			lp++
		} else {
			res = append(res, vals[rp])
			rp++
		}
	}

	if lp > mid {
		res = append(res, vals[rp:right+1]...)
	} else {
		res = append(res, vals[lp:mid+1]...)
	}

	idx := 0
	for i := left; i <= right ; i++ {
		vals[i] = res[idx]
		idx++
	}
}