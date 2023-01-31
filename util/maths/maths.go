package maths

func Max(vals ...int) int {
	res := vals[0]

	for _, v := range vals[1:] {
		if v > res {
			res = v
		}
	}

	return res
}

func Min(vals ...int) int {
	res := vals[0]

	for _, v := range vals[1:] {
		if v < res {
			res = v
		}
	}

	return res
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
