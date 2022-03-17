package compare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	x, y := 1, 2
	res := Max(x, y)

	assert.Equal(t, 2, res)
}
