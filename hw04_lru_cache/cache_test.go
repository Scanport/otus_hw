package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)

		wasInCache = c.Set("ccc", 400)
		require.False(t, wasInCache)

		val, ok = c.Get("ccc")
		require.True(t, ok)
		require.Equal(t, 400, val)

		c.Clear()

		_, ok = c.Get("ccc")
		require.False(t, ok)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(3)

		c.Set("a", 100)
		c.Set("b", 200)
		c.Set("c", 300)

		// pushed out a
		c.Set("d", 400)
		// ["b", "c", "d"]
		_, ok := c.Get("a")
		require.False(t, ok)

		_, ok = c.Get("c")
		require.True(t, ok)
		// ["b", "d", "c"]

		// pushed out b
		c.Set("e", 500)
		// ["d", "c", "e"]
		_, ok = c.Get("b")
		require.False(t, ok)

		// pushed out e
		c.Get("c")
		// ["d", "e", "c"]
		c.Set("d", 500)
		// ["e", "c", "d"]
		c.Set("f", 600)
		// ["c", "d", "f"]
		_, ok = c.Get("e")
		require.False(t, ok)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
