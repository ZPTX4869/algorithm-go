package binarys

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hammingWeight(t *testing.T) {
	var num uint32
	var res int

	num = 11
	res = hammingWeight2(num)
	t.Run("case1", func(t *testing.T) {
		assert.Equal(t, 3, res)
	})
}
