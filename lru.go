package lru

import (
	"container/list"
	"sync"
)

type LRUCache interface {
	// Add adds new value to the cache by key (with the highest priority).
	// In case of duplicate key it will replace old value to new value and set the highest priority.
	// In case of oversize of cache capacity it will remove the lowest priority element.
	Add(key string, value any)
	// Get Returns element by key with the boolean flag indicating the element is in the cache
	// If the element is in the cache, sets the highest priority to this element
	Get(key string) (value any, found bool)
}

type cache struct {
	capacity int
	queue    *list.List
	items    map[string]*list.Element
	mu       sync.Mutex
}

func NewLRUCache(capacity int) *cache {
	return &cache{capacity: capacity, queue: list.New(), items: make(map[string]*list.Element)}
}

type item struct {
	key   string
	value any
}

func (c *cache) Add(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if element, found := c.items[key]; found {
		element.Value.(*item).value = value
		c.queue.MoveToFront(element)
		return
	}

	if c.queue.Len() == c.capacity {
		c.removeBackElement()
	}

	item := &item{key: key, value: value}
	elem := c.queue.PushFront(item)
	c.items[key] = elem
}

func (c *cache) Get(key string) (value any, found bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if element, found := c.items[key]; found {
		c.queue.MoveToFront(element)
		return element.Value.(*item).value, true
	}

	return nil, false
}

func (c *cache) removeBackElement() {
	if element := c.queue.Back(); element != nil {
		item := c.queue.Remove(element).(*item)
		delete(c.items, item.key)
	}
}
