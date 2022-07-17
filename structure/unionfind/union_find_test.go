package unionfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var uf UnionFind

func init() {
	uf = New(10)
	uf.Union(0, 1)
	uf.Union(0, 6)
	uf.Union(0, 3)
	uf.Union(2, 5)
}

func TestUnionFind_Union(t *testing.T) {
	uf.Union(5, 1)
	assert.Equal(t, 5, uf.Count())
}

func TestUnionFind_Find(t *testing.T) {
	uf.Union(5, 1)
	root := uf.Find(1)
	assert.Equal(t, 2, root)
}

func TestUnionFind_Connected(t *testing.T) {
	uf.Union(5, 1)
	res := uf.Connected(1, 5)
	assert.True(t, res)
}
