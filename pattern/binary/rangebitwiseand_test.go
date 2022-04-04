package binary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rangeBitwiseAnd(t *testing.T) {
	res := rangeBitwiseAnd(1, 2)
	assert.Equal(t, 0, res)
}
