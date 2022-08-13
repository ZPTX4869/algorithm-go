package maths

import "golang.org/x/exp/constraints"

func Sum[T constraints.Signed](vals []T) T {
	var sum T

	for i := 0; i < len(vals); i++ {
		sum += vals[i]
	}

	return sum
}

func Abs[T constraints.Signed](x T) T {
    if x < 0 {
        return -x
    }

    return x
}

func Max[T constraints.Ordered](vals ...T) T {
	res := vals[0]

	for _, v := range vals[1:] {
		if v > res {
			res = v
		}
	}
	
	return res
}

func Min[T constraints.Ordered](vals ...T) T {
	res := vals[0]

	for _, v := range vals[1:] {
		if v < res {
			res = v
		}
	}
	
	return res
}