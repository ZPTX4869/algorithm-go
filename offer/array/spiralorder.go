package array

type border struct {
	top, bottom int
	left, right int
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	res := make([]int, 0)
	m, n := len(matrix), len(matrix[0])
	b := border{
		0, m - 1,
		0, n - 1,
	}

	for {
		for j := b.left; j <= b.right; j++ {
			res = append(res, matrix[b.top][j])
		}
		if len(res) == m*n {
			break
		}
		b.top += 1

		for i := b.top; i <= b.bottom; i++ {
			res = append(res, matrix[i][b.right])
		}
		if len(res) == m*n {
			break
		}
		b.right -= 1

		for j := b.right; j >= b.left; j-- {
			res = append(res, matrix[b.bottom][j])
		}
		if len(res) == m*n {
			break
		}
		b.bottom -= 1

		for i := b.bottom; i >= b.top; i-- {
			res = append(res, matrix[i][b.left])
		}
		if len(res) == m*n {
			break
		}
		b.left += 1
	}

	return res
}
