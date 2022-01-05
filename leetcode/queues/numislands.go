package queues

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	islandCnt := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				islandCnt++
				bfs(grid, i, j)
			}
		}
	}

	return islandCnt
}

func bfs(grid [][]byte, row, col int) {
	grid[row][col] = '0'

	numRows := len(grid)
	numCols := len(grid[0])

	queue := []int{row*numCols + col}
	for len(queue) > 0 {
		i, j := queue[0]/numCols, queue[0]%numCols
		queue = queue[1:]

		if i-1 >= 0 && grid[i-1][j] == '1' {
			grid[i-1][j] = '0'
			queue = append(queue, (i-1)*numCols+j)
		}
		if i+1 < numRows && grid[i+1][j] == '1' {
			grid[i+1][j] = '0'
			queue = append(queue, (i+1)*numCols+j)
		}

		if j-1 >= 0 && grid[i][j-1] == '1' {
			grid[i][j-1] = '0'
			queue = append(queue, i*numCols+(j-1))
		}
		if j+1 < numCols && grid[i][j+1] == '1' {
			grid[i][j+1] = '0'
			queue = append(queue, i*numCols+(j+1))
		}
	}
}

//func numIslands(grid [][]byte) int {
//	if len(grid) == 0 {
//		return 0
//	}
//
//	islandCnt := 0
//	for i := 0; i < len(grid); i++ {
//		for j := 0; j < len(grid[0]); j++ {
//			if grid[i][j] == '1' {
//				islandCnt++
//				dfs(grid, i, j)
//			}
//		}
//	}
//
//	return islandCnt
//}
//
//func dfs(grid [][]byte, i, j int) {
//	grid[i][j] = '0'
//
//	numRows := len(grid)
//	numCols := len(grid[0])
//
//	if i-1 >= 0 && grid[i-1][j] == '1' {
//		dfs(grid, i-1, j)
//	}
//	if i+1 < numRows && grid[i+1][j] == '1' {
//		dfs(grid, i+1, j)
//	}
//
//	if j-1 >= 0 && grid[i][j-1] == '1' {
//		dfs(grid, i, j-1)
//	}
//	if j+1 < numCols && grid[i][j+1] == '1' {
//		dfs(grid, i, j+1)
//	}
//}
