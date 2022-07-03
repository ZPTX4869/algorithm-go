package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	assert.Equal(t, cache.Get(1), 1)

	cache.Put(3, 3)
	assert.Equal(t,cache.Get(2), -1)

	cache.Put(4, 4)
	assert.Equal(t,cache.Get(1), -1)
	assert.Equal(t,cache.Get(3), 3)
	assert.Equal(t,cache.Get(4), 4)
}

func TestLRUCache2(t *testing.T) {
	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(1, 1)
	cache.Put(2, 3)
	cache.Put(4, 1)

	assert.Equal(t,cache.Get(1), -1)
	assert.Equal(t,cache.Get(2), 3)
}
