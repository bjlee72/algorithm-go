package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLruAdd(t *testing.T) {
	cache := NewLRUCache(3)
	cache.Put("a", testValue{1})
	cache.Put("b", testValue{2})
	cache.Put("c", testValue{3})
	cache.Put("d", testValue{4})
	// to test the update of the existing item.
	// and this shouldn't affect the current size of the cache.
	cache.Put("c", testValue{7})

	assert.Equal(t, 3, cache.Length())
	assert.Nil(t, cache.Get("a"))

	assert.Equal(t, "2", cache.Get("b").String())
	assert.Equal(t, "7", cache.Get("c").String())
	assert.Equal(t, "4", cache.Get("d").String())
}
