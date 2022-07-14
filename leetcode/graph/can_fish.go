package graph

var hasCycle bool
var onPath []bool
var visited []bool

func canFinish(numCourses int, prerequisites [][]int) bool {
	hasCycle = false
	onPath = make([]bool, numCourses)
	visited = make([]bool, numCourses)

	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = []int{}
	}

	for i := 0; i < len(prerequisites); i++ {
		from, to := prerequisites[i][1], prerequisites[i][0]
		graph[from] = append(graph[from], to)
	}

	for i := 0; i < numCourses; i++ {
		traverse(graph, i)
	}

	return !hasCycle
}

func traverse(graph [][]int, s int) {
	if onPath[s] == true {
		hasCycle = true
		return
	}

	if visited[s] == true {
		return
	}

	onPath[s] = true
	visited[s] = true
	for i := 0; i < len(graph[s]); i++ {
		traverse(graph, graph[s][i])
	}
	onPath[s] = false
}
