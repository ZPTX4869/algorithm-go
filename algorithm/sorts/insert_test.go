package sorts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertSort_int(t *testing.T) {
	vals := []int{5, 4, 3, 2, 1}
	InsertSort(vals)

	assert.Equal(t, []int{1, 2, 3, 4, 5}, vals)
}

func TestInsertSort_float(t *testing.T) {
	vals := []int{3.0, 2.0, 1.0}
	InsertSort(vals)

	assert.Equal(t, []int{1.0, 2.0, 3.0}, vals)
}

func TestInsertSort_byte(t *testing.T) {
	vals := []int{'c', 'b', 'a'}
	InsertSort(vals)

	assert.Equal(t, []int{'a', 'b', 'c'}, vals)
}
