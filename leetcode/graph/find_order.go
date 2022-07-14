package graph

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = []int{}
	}
	for i := 0; i < len(prerequisites); i++ {
		from, to := prerequisites[i][1], prerequisites[i][0]
		graph[from] = append(graph[from], to)
	}

	inDegree := make([]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		to := prerequisites[i][0]
		inDegree[to] += 1
	}

	queue := []int{}
	for vertex, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, vertex)
		}
	}

	var cnt int
	path := []int{}
	for len(queue) > 0 {
		cur := queue[0]
		cnt += 1
		path = append(path, cur)
		for _, next := range graph[cur] {
			inDegree[next] -= 1
			// 当一个节点的入度为0，说明前置条件全部满足
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
		queue = queue[1:]
	}

	if cnt != numCourses {
		return []int{}
	}

	return path
}
