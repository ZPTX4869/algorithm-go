package slices

import (
	"errors"

	"golang.org/x/exp/constraints"
)

func Sum[T constraints.Signed](vals []T) T {
	var sum T

	for i := 0; i < len(vals); i++ {
		sum += vals[i]
	}

	return sum
}

func Fill[T any](elems []T, val T) error {
	if elems == nil {
		return errors.New("nil slice can not be filled")
	}

	for i := range elems {
		elems[i] = val
	}

	return nil
}

func Filter[T comparable](elems []T, predicate func(x T) bool) []T {
	res := make([]T, 0)

	for _, val := range elems {
		if predicate(val) {
			res = append(res, val)
		}
	}

	return res
}
