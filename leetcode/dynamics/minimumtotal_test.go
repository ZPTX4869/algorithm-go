package dynamics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minimumTotal(t *testing.T) {
	//triangle := [][]int{
	//	{2},
	//	{3, 4},
	//	{6, 5, 7},
	//	{4, 1, 8, 3},
	//}

	triangle := [][]int{
		{1},
		{2, 3},
	}

	assert.Equal(t, 11, minimumTotal(triangle))
}
