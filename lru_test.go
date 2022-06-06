package lru

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAdd(t *testing.T) {
	cache := NewLRUCache(3)

	cache.Add("k1", "v1")
	cache.Add("k2", "v2")
	cache.Add("k3", "v3")
	// queue: k3 -> k2 -> k1

	v, found := cache.Get("k1")
	require.True(t, found)
	require.Equal(t, "v1", v)
	// queue: k1 -> k3 -> k2

	v, found = cache.Get("k2")
	require.True(t, found)
	require.Equal(t, "v2", v)
	// queue: k2 -> k1 -> k3

	v, found = cache.Get("k3")
	require.True(t, found)
	require.Equal(t, "v3", v)
	// queue: k3 -> k2 -> k1

	cache.Add("k4", "v4")
	// queue: k4 -> k3 -> k2

	_, found = cache.Get("k1")
	require.False(t, found)
	// queue: k4 -> k3 -> k2

	v, found = cache.Get("k4")
	require.True(t, found)
	require.Equal(t, "v4", v)
	// queue: k4 -> k3 -> k2

	cache.Add("k5", "v5")
	// queue: k5 -> k4 -> k3
	_, found = cache.Get("k2")
	require.False(t, found)
	// queue: k5 -> k4 -> k3

	_, found = cache.Get("k3")
	require.True(t, found)
	// queue: k3 -> k5 -> k4

	cache.Add("k4", "v4_new")
	// queue: k4 -> k3 -> k5
	v, found = cache.Get("k4")
	require.True(t, found)
	require.Equal(t, "v4_new", v)
}
