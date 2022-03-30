package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_integerBreak(t *testing.T) {
	var n int
	var res int
	
	n = 10
	res = integerBreak(n)
	t.Run("case1", func(t *testing.T) {
		assert.Equal(t, 36, res)
	})
}
