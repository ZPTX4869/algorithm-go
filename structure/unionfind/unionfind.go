package unionfind

type UnionFind struct {
	// 联通分量
	count int
	// 父节点数组
	parent []int
}

func New(numNodes int) UnionFind {
	parent := make([]int, numNodes)
	for i := range parent {
		parent[i] = i
	}

	return UnionFind{
		count:  numNodes,
		parent: parent,
	}
}

func (uf *UnionFind) Union(p, q int) {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)

	if rootP == rootQ {
		return
	}

	uf.parent[rootQ] = rootP
	uf.count -= 1
}

func (uf *UnionFind) Find(x int) int {
	if x != uf.parent[x] {
		uf.parent[x] = uf.Find(uf.parent[x])
	}

	return uf.parent[x]
}

func (uf *UnionFind) Connected(p, q int) bool {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)

	return rootP == rootQ
}

func (uf UnionFind) Count() int {
	return uf.count
}
