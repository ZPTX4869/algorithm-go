package dynamic

import "fmt"

func backPack(m int, A []int) int {
	n := len(A)
	f := make([][]bool, n+1)
	for i := range f {
		f[i] = make([]bool, m+1)
	}

	f[0][0] = true
	for i := 1; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			f[i][j] = f[i-1][j]
			if j-A[i-1] >= 0 && f[i-1][j-A[i-1]] {
				f[i][j] = true
			}
		}
		fmt.Printf("%d: %v\n", i, f[i][1:])
	}

	res := 0
	for i := m; i >= 0; i-- {
		if f[n][i] {
			res = i
			break
		}
	}

	return res
}

func backPackII(m int, A []int, V []int) int {
	getMax := func (x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	n := len(A)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
		for j := 0; j < m+1; j++ {
			f[i][j] = -1
		}
	}

	f[0][0] = 0
	for i := 1; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			f[i][j] = f[i-1][j]
			if j-A[i-1] >= 0 && f[i-1][j-A[i-1]] >= 0 {
				f[i][j] = getMax(f[i-1][j], f[i-1][j-A[i-1]] + V[i-1])
			}
		}
	}

	max := 0
	for i := m; i >= 0; i-- {
		if f[n][i] > max {
			max = f[n][i]
		}
	}

	return max
}
