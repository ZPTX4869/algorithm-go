package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_constructArr(t *testing.T) {
	var a []int
	var res []int

	a = []int{1, 2, 3, 4, 5}
	res = constructArr(a)
	t.Run("case1", func(t *testing.T) {
		assert.Equal(t, []int{120, 60, 40, 30, 24}, res)
	})
}
