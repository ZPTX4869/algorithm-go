package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRU(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		cache := NewLRU(1)
		cache.Put(2, 1)
		assert.Equal(t, 1, cache.Get(2))
		cache.Put(3, 1)
		assert.Equal(t, 1, cache.Get(3))
	})

	t.Run("case2", func(t *testing.T) {
		cache := NewLRU(2)
		assert.Equal(t, -1, cache.Get(2))
		cache.Put(2, 6)
		assert.Equal(t, -1, cache.Get(1))
		cache.Put(1, 5)
		cache.Put(1, 2)
		assert.Equal(t, 2, cache.Get(1))
		assert.Equal(t, 6, cache.Get(2))
	})
}
