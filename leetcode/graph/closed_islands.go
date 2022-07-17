package graph

// 统计封闭岛屿的数目：https://leetcode.cn/problems/number-of-closed-islands/
func closedIsland(grid [][]int) int {
	var res int
	m, n := len(grid), len(grid[0])

	// 淹掉上下两个边界关联的岛屿
	for i := 0; i < m; i++ {
		closedIsland_dfs(grid, i, 0)
		closedIsland_dfs(grid, i, n-1)
	}

	// 淹掉左右两个边界关联的岛屿
	for j := 0; j < n; j++ {
		closedIsland_dfs(grid, 0, j)
		closedIsland_dfs(grid, m-1, j)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				res += 1
				closedIsland_dfs(grid, i, j)
			}
		}
	}

	return res
}

func closedIsland_dfs(grid [][]int, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}

	if grid[i][j] == 1 {
		return
	}

	grid[i][j] = 1
	closedIsland_dfs(grid, i-1, j)
	closedIsland_dfs(grid, i+1, j)
	closedIsland_dfs(grid, i, j-1)
	closedIsland_dfs(grid, i, j+1)
}
