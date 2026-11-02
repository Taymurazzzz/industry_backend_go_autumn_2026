package main

type Cache[K comparable, V any] struct {
	capacity int
	items    map[K]V
}

func NewCache[K comparable, V any](capacity int) *Cache[K, V] {
	return &Cache[K, V]{capacity: capacity, items: make(map[K]V)}
}
func (c *Cache[K, V]) Get(k K) (v V, ok bool) {
	value, ok := c.items[k]
	return value, ok
}
func (c *Cache[K, V]) Set(k K, v V) bool {
	capacity := c.capacity
	_, ok := c.items[k]
	if capacity <= 0 || len(c.items) >= capacity && !ok {
		return false
	}
	c.items[k] = v
	return true
}
