package binary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_add(t *testing.T) {
	var a, b int
	var res int

	a, b = 111, 222
	res = add(a, b)
	t.Run("case1", func(t *testing.T) {
		assert.Equal(t, 333, res)
	})

	a, b = 4, -2
	res = add(a, b)
	t.Run("case2", func(t *testing.T) {
		assert.Equal(t, 2, res)
	})

	a, b = 1, -2
	res = add(a, b)
	t.Run("case3", func(t *testing.T) {
		assert.Equal(t, -1, res)
	})
}
