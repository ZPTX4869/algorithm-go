package slices

import "errors"

func Fill[T any](elems []T, val T) error {
	if elems == nil {
		return errors.New("nil slice can not be filled")
	}

	for i := range elems {
		elems[i] = val
	}

	return nil
}

func FilterSlice[T comparable](elems []T, predicate func(x T) bool) []T {
	res := make([]T, 0)

	for _, val := range elems {
		if predicate(val) == true {
			res = append(res, val)
		}
	}

	return res
}
