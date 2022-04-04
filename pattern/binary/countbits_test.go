package binary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countBits(t *testing.T) {
	res := countBits(2)
	assert.Equal(t, []int{0, 1, 1}, res)
}
