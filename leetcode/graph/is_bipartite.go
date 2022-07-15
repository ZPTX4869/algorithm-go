package graph

// 判断是否为二分图：https://leetcode.cn/problems/is-graph-bipartite/
func isBipartite(graph [][]int) bool {
	res := true
	color := make([]bool, len(graph))
	visited := make([]bool, len(graph))

	// DFS
	var dfs func(root int)
	dfs = func(root int) {
		if res == false {
			return
		}

		visited[root] = true
		for _, neighbor := range graph[root] {
			if visited[neighbor] == false {
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
		if visited[i] == false {
			dfs(i)
		}
	}

	return res
}

func isBipartite_(graph [][]int) bool {
	res := true
	color := make([]bool, len(graph))
	visited := make([]bool, len(graph))

	// BFS
	var bfs func(root int)
	bfs = func(root int) {
		queue := []int{root}
		visited[root] = true

		for len(queue) > 0 {
			cur := queue[0]

			for _, neighbor := range graph[cur] {
				if visited[neighbor] == false {
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
		if visited[i] == false {
			bfs(i)
		}
	}

	return res
}
