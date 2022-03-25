package binarys

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myPow(t *testing.T) {
	var x float64
	var n int
	var res float64

	x, n = 5, 4
	res = myPow(x, n)
	assert.Equal(t, 625.00000, res)

	x, n = 2, -2
	res = myPow(x, n)
	assert.Equal(t, 0.25000, res)
}
