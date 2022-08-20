package maths

import "golang.org/x/exp/constraints"

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

func Abs[T constraints.Signed](x T) T {
    if x < 0 {
        return -x
    }

    return x
}
