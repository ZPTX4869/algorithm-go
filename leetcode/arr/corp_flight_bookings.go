package arr

type diff []int

func (d diff) increment(i, j int, val int) {
	d[i-1] += val
	if j < len(d) {
		d[j] -= val
	}
}

func (d diff) restore() []int {
	res := make([]int, len(d))
	res[0] = d[0]

	for i := 1; i < len(d); i++ {
		res[i] = res[i-1] + d[i]
	}

	return res
}

func corpFlightBookings(bookings [][]int, n int) []int {
	var d diff = make([]int, n)

	for i := 0; i < len(bookings); i++ {
		d.increment(bookings[i][0], bookings[i][1], bookings[i][2])
	}

	return d.restore()
}
