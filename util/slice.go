package util

func FilterSlice[T comparable](vals []T, predicate func(x T) bool) []T {
	res := make([]T, 0)

	for _, val :=  range vals {
		if predicate(val) == true {
			res = append(res, val)
		}
	}

	return res
}
