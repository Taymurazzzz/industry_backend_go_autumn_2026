package main

import "container/list"

type entry[K comparable, V any] struct {
	key   K
	value V
}
type LRUCache[K comparable, V any] struct {
	capacity int
	ll       list.List
	items    map[K]*list.Element
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{capacity: capacity, ll: list.List{}, items: make(map[K]*list.Element)}
}
func (c *LRUCache[K, V]) Get(key K) (value V, ok bool) {
	item, ok := c.items[key]
	if !ok {
		return value, false
	}
	c.ll.MoveToFront(item)
	return item.Value.(*entry[K, V]).value, true
}
func (c *LRUCache[K, V]) Set(key K, value V) {
	if c.capacity <= 0 {
		return
	}
	if item, ok := c.items[key]; ok {
		item.Value.(*entry[K, V]).value = value
		return
	}
	if len(c.items) >= c.capacity {
		last := c.ll.Back()
		if last == nil {
			panic("LRU: cache is full but list is empty")
		}

		record := last.Value.(*entry[K, V])
		c.ll.Remove(last)
		delete(c.items, record.key)
	}
	item := c.ll.PushFront(&entry[K, V]{
		key:   key,
		value: value,
	})
	c.items[key] = item
}

type LRU[K comparable, V any] interface {
	Get(K) (V, bool)
	Set(K, V)
}
