package graph

import (
	"algorithm-go/structure/unionfind"
	"algorithm-go/util"
	"sort"
)

// Min Cost to Connect All Points: https://leetcode.cn/problems/min-cost-to-connect-all-points/
func minCostConnectPoints(points [][]int) int {
	uf := unionfind.New(len(points))
	edges := make([][]int, 0)

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dis := manhattan(points[i], points[j])
			edges = append(edges, []int{i, j, dis})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	var mst int
	for _, edge := range edges {
		p1, p2, dis := edge[0], edge[1], edge[2]

		if uf.Connected(p1, p2) {
			continue
		}

		uf.Union(p1, p2)
		mst += dis
	}

	return mst
}

func manhattan(p1, p2 []int) int {
	return util.Abs(p1[0]-p2[0]) + util.Abs(p1[1]-p2[1])
}
