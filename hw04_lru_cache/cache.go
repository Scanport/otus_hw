package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if item, exists := c.items[key]; exists {
		// Update existing item
		item.Value = value
		c.queue.MoveToFront(item)
		return true
	}
	// Add new item
	item := c.queue.PushFront(value)
	c.items[key] = item

	// Remove the least recently used item if cache is full
	if c.queue.Len() > c.capacity {
		oldest := c.queue.Back()
		delKey := findKeysByValue(c.items, oldest)
		delete(c.items, delKey)
		c.queue.Remove(oldest)
	}
	return false
}

func findKeysByValue(m map[Key]*ListItem, val *ListItem) Key {
	for key, item := range m {
		if item == val {
			return key
		}
	}
	return ""
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if item, exists := c.items[key]; exists {
		c.queue.MoveToFront(item)
		return item.Value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem)
}
