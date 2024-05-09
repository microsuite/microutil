package lru

import "container/list"

// entry is used to hold a value in the cache.
type entry struct {
	key, value interface{}
}

// Cache is an LRU cache. It is not safe for concurrent access.
type Cache struct {
	// size is the maximum capacity of the cache. 0 means no limit.
	size int

	// list is the doubly-linked list that stores the cache entries.
	list *list.List

	// cache is a mapping of keys to entries.
	cache map[interface{}]*list.Element
}

// removeElement removes an element from the cache.
func (c *Cache) removeElement(elem *list.Element) {
	c.list.Remove(elem)
	delete(c.cache, elem.Value.(*entry).key)
}

// removeOldest removes the oldest item from the cache.
func (c *Cache) removeOldest() {
	elem := c.list.Back()
	if elem != nil {
		c.removeElement(elem)
	}
}

// Get looks up a key's value from the cache.
func (c *Cache) Get(key interface{}) interface{} {
	if c.cache == nil {
		return nil
	}

	// Look up the element for the key.
	if elem, ok := c.cache[key]; ok {
		// Move the element to the front of the list.
		c.list.MoveToFront(elem)
		return elem.Value.(*entry).value
	}
	return nil
}

// Add adds a value to the cache.
func (c *Cache) Add(key, value interface{}) {
	if c.cache == nil {
		c.cache = make(map[interface{}]*list.Element)
		c.list = list.New()
	}

	// Check if the key already exists in the cache.
	if elem, ok := c.cache[key]; ok {
		// Move the existing element to the front of the list.
		c.list.MoveToFront(elem)
		elem.Value.(*entry).value = value
	}

	// Add a new element to the front of the list.
	elem := c.list.PushFront(&entry{key, value})
	c.cache[key] = elem

	// Remove the oldest element if the cache is at capacity.
	if c.size != 0 && c.list.Len() > c.size {
		c.removeOldest()
	}
}

// Remove removes the provided key from the cache.
func (c *Cache) Remove(key interface{}) {
	if c.cache == nil {
		return
	}

	if elem, ok := c.cache[key]; ok {
		c.removeElement(elem)
	}
}

// Clear removes all key-value pairs from the cache.
func (c *Cache) Clear() {
	if c.cache == nil {
		return
	}

	for _, elem := range c.cache {
		c.removeElement(elem)
	}

	c.list = list.New()
	c.cache = make(map[interface{}]*list.Element)
}

// Contains returns true if the cache contains the key, otherwise returns false..
func (c *Cache) Contains(key interface{}) bool {
	if c.cache == nil {
		return false
	}

	// Look up the element for the key.
	if elem, ok := c.cache[key]; ok {
		// Move the element to the front of the list.
		c.list.MoveToFront(elem)
		return true
	}
	return false
}

// Keys returns a slice of the keys in the cache.
func (c *Cache) Keys() []interface{} {
	if c.cache == nil {
		return nil
	}

	keys := make([]interface{}, 0, c.Len())
	for _, elem := range c.cache {
		keys = append(keys, elem.Value.(*entry).key)
	}

	return keys
}

// Len returns the number of items in the cache.
func (c *Cache) Len() int {
	if c.cache == nil {
		return 0
	}

	return len(c.cache)
}

// New creates a new LRU cache.
func New(size int) *Cache {
	return &Cache{
		size:  size,
		list:  list.New(),
		cache: make(map[interface{}]*list.Element),
	}
}
