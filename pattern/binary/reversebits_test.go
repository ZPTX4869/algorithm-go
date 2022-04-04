package binary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseBits(t *testing.T) {
	num := uint32(1)
	res := reverseBits(num)
	assert.Equal(t, uint32(2147483648), res)
}
