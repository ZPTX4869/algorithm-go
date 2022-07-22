package graph

import (
	"fmt"
	"testing"
)

func TestDijkstra(t *testing.T) {
	edges := [][]int{
		{1, 0, 1},
		{1, 2, 1},
		{2, 3, 1},
	}
	graph := make([][][]int, 4)
	for i := range graph {
		graph[i] = make([][]int, 0)
	}

	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		graph[from] = append(graph[from], []int{to, weight})
	}

	distTo := Dijkstra(1, graph)
	fmt.Println(distTo)
}
