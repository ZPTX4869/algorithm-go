package sort

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

func QuickSort[T constraints.Ordered](vals []T) {
	var sort func(left, right int)
	sort = func(start, end int) {
		if start >= end {
			return
		}

		pivot := vals[start]
		left, right := start, end
		for left < right {
			for left < right && vals[right] >= pivot {
				right--
			}
			for left < right && vals[left] <= pivot {
				left++
			}
			vals[left], vals[right] = vals[right], vals[left]
		}
		vals[start], vals[left] = vals[left], vals[start]

		sort(start, left-1)
		sort(left+1, end)
	}

	rand.Shuffle(len(vals), func(i, j int) {
		vals[i], vals[j] = vals[j], vals[i]
	})
	
	sort(0, len(vals)-1)
}
