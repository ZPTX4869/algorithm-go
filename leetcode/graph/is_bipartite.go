package graph

// 判断是否为二分图：https://leetcode.cn/problems/is-graph-bipartite/
func isBipartite(graph [][]int) bool {
	res := true
	color := make([]bool, len(graph))
	visited := make([]bool, len(graph))

	// DFS
	var dfs func(root int)
	dfs = func(root int) {
		if !res {
			return
		}

		visited[root] = true
		for _, neighbor := range graph[root] {
			if !visited[neighbor] {
				color[neighbor] = !color[root]
				dfs(neighbor)
			} else {
				if color[root] == color[neighbor] {
					res = false
				}
			}
		}
	}

	for i := 0; i < len(graph); i++ {
		if !visited[i] {
			dfs(i)
		}
	}

	return res
}

func isBipartite2(graph [][]int) bool {
	res := true
	color := make([]bool, len(graph))
	visited := make([]bool, len(graph))

	// BFS
	bfs := func(root int) {
		queue := []int{root}
		visited[root] = true

		for len(queue) > 0 {
			cur := queue[0]

			for _, neighbor := range graph[cur] {
				if !visited[neighbor] {
					color[neighbor] = !color[cur]
					visited[neighbor] = true
					queue = append(queue, neighbor)
				} else {
					if color[neighbor] == color[cur] {
						res = false
						return
					}
				}
			}

			queue = queue[1:]
		}
	}

	for i := 0; i < len(graph); i++ {
		if !visited[i] {
			bfs(i)
		}
	}

	return res
}
