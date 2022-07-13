package graph

// 所有可能的路径：https://leetcode.cn/problems/all-paths-from-source-to-target/
func allPathsSourceTarget(graph [][]int) [][]int {
    res := [][]int{}
    path := []int{}

	var traverse func(idx int)
	traverse = func(idx int) {
		path = append(path, idx)
	
		n := len(graph)
		if idx == n-1 {
			res = append(res, append([]int{}, path...))
		}
	
		for _, neighbor := range graph[idx] {
			traverse(neighbor)
		}
	
		path = path[:len(path)-1]
	}

    traverse(0)

    return res
}
