package searchs

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	n, m := len(matrix), len(matrix[0])
	binarySearch := func(arr []int) bool {
		start, end := 0, m

		for start != end {
			mid := (end-start)/2 + start

			if target < arr[mid] {
				end = mid
			} else if target > arr[mid] {
				start = mid + 1
			} else {
				return true
			}
		}

		return false
	}

	for i := 0; i < n; i++ {
		if target < matrix[i][0] || target > matrix[i][m-1] {
			continue
		}

		if binarySearch(matrix[i]) {
			return true
		}
	}

	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	n, m := len(matrix), len(matrix[0])

	i, j := 0, m-1
	for i < n && j >= 0 {
		if target < matrix[i][j] {
			j--
		} else if target > matrix[i][j] {
			i++
		} else {
			return true
		}
	}

	return false
}