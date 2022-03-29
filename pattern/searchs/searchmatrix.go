package searchs

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	left, right := 0, to1D(&matrix, len(matrix)-1, len(matrix[0]))

	for left < right {
		mid := left + (right-left)/2
		mRow, mCol := to2D(&matrix, mid)

		if matrix[mRow][mCol] == target {
			return true
		} else if matrix[mRow][mCol] < target {
			left = mid + 1
		} else if matrix[mRow][mCol] > target {
			right = mid
		}
	}

	return false
}

func to1D(mat *[][]int, row, col int) int {
	numCols := len((*mat)[0])

	return row*numCols + col
}

func to2D(mat *[][]int, idx int) (rowNum, colNum int) {
	numCols := len((*mat)[0])

	rowNum = idx / numCols
	colNum = idx % numCols

	return
}
