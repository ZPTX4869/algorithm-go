package graph

import (
	"algorithm-go/util/slices"
	"math"
)

type State struct {
	nodeID        int
	distFromStart int
}

func Dijkstra(start int, graph [][][]int) []int {
	numNodes := len(graph)
	distTo := make([]int, numNodes)
	queue := []*State{{start, 0}}

	err := slices.Fill(distTo, math.MaxInt)
	if err != nil {
		panic(err)
	}
	distTo[start] = 0

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		curID := state.nodeID
		curFromStart := state.distFromStart

		if curFromStart > distTo[curID] {
			continue
		}

		for _, nbrInfo := range graph[curID] {
			nbrID := nbrInfo[0]
			nbrFromStart := nbrInfo[1] + distTo[curID]

			if nbrFromStart < distTo[nbrID] {
				distTo[nbrID] = nbrFromStart
				queue = append(queue, &State{nbrID, nbrFromStart})
			}
		}
	}

	return distTo
}
