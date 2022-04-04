package binary

func myPow(x float64, n int) float64 {
	var res float64
	res = 1

	if n < 0 {
		x = 1/x
		n = -n
	}

	for n != 0 {
		b := n & 1

		if b > 0 {
			res *= x
		}

		n = n >> 1
		x *= x
	}

	return res
}
