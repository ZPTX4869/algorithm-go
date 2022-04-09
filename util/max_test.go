package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	var x, y int
	var res int

	t.Run("case1", func(t *testing.T) {
		x, y = 2, 3
		res = Max(x, y)
		assert.Equal(t, 3, res)
	})
}
