package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_updateMatrix(t *testing.T) {
	mat := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}

	distMat := updateMatrix(mat)

	assert.Equal(t, [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 2, 1},
	}, distMat)
}
