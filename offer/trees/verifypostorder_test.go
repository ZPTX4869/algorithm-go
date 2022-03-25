package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_verifyPostorder(t *testing.T) {
	var postorder []int
	var res bool

	postorder = []int{1, 6, 3, 2, 5}
	res = verifyPostorder(postorder)
	t.Run("case1", func(t *testing.T) {
		assert.Equal(t, false, res)
	})

	postorder = []int{1, 3, 2, 6, 5}
	res = verifyPostorder(postorder)
	t.Run("case2", func(t *testing.T) {
		assert.Equal(t, true, res)
	})

	postorder = []int{1, 2, 5, 10, 6, 9, 4, 3}
	res = verifyPostorder(postorder)
	t.Run("case3", func(t *testing.T) {
		assert.Equal(t, false, res)
	})
}
