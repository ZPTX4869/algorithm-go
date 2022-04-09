package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	var vals []int

	t.Run("case1", func(t *testing.T) {
		vals = []int{5, 4, 3, 2, 1}
		QuickSort(vals)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, vals)
	})

	t.Run("case2", func(t *testing.T) {
		vals = []int{2, 11, 4, 7}
		QuickSort(vals)
		assert.Equal(t, []int{2, 4, 7, 11}, vals)
	})
}
