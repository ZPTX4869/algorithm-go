package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_firstBadVersion(t *testing.T) {
	n, bad := 5, 4

	setBadVersion(bad)
	res := firstBadVersion(n)

	assert.Equal(t, bad, res)
}
