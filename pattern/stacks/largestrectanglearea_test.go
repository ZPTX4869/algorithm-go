package stacks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_largestRectangleArea(t *testing.T) {
	heights := []int{3, 6, 5, 7, 4, 8, 1, 0}
	maxArea := largestRectangleArea(heights)

	assert.Equal(t, 20, maxArea)
}
