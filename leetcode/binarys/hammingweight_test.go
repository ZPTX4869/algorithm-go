package binarys

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hammingWeight(t *testing.T) {
	var num uint32
	num = 4
	assert.Equal(t, 1, hammingWeight(uint32(num)))
}
