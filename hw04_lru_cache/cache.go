package hw04lrucache

type Key string

type data struct {
	k Key
	v interface{}
}

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
		item.Value = data{key, value}
		c.queue.MoveToFront(item)
		return true
	}
	// Add new item
	item := c.queue.PushFront(data{key, value})
	c.items[key] = item

	// Remove the least recently used item if cache is full
	if c.queue.Len() > c.capacity {
		oldest := c.queue.Back()
		delete(c.items, oldest.Value.(data).k)
		c.queue.Remove(oldest)
	}
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if item, exists := c.items[key]; exists {
		c.queue.MoveToFront(item)
		return item.Value.(data).v, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem)
}
