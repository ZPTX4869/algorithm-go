package dynamics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_uniquePaths(t *testing.T) {
	m, n := 3, 7
	assert.Equal(t, 28, uniquePaths(m, n))
}

func Test_uniquePathsWithObstacles(t *testing.T) {
	obstacleGrid := [][]int{
		{0, 0},
		{1, 1},
		{0, 0},
	}

	assert.Equal(t, 0, uniquePathsWithObstacles(obstacleGrid))
}
