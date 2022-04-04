package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numIslands(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	res := numIslands(grid)

	assert.Equal(t, 1, res)
}
