package practices

import "fmt"

type Node struct {
	prev  *Node
	next  *Node
	key   string
	value string
}

type LRUCache struct {
	head     *Node
	tail     *Node
	capacity int
	cache    map[string]*Node
}

func NewLRUCache(capacity int) *LRUCache {
	lru := &LRUCache{
		head:     &Node{},
		tail:     &Node{},
		capacity: capacity,
		cache:    make(map[string]*Node),
	}

	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (lru *LRUCache) Set(key string, value string) {
	node, ok := lru.cache[key]
	if !ok {
		node = &Node{
			key:   key,
			value: value,
		}
		lru.cache[key] = node
	}

	// remove node from current position
	if (node.prev != nil) || (node.next != nil) {
		node.prev.next = node.next
		node.next.prev = node.prev
	}

	// put node to head
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node

	// check length
	if len(lru.cache) > lru.capacity {
		last_node := lru.tail.prev

		last_node.prev.next = last_node.next
		last_node.next.prev = last_node.prev

		delete(lru.cache, last_node.key)

	}
}

func (lru *LRUCache) Get(key string) (string, bool) {
	node, ok := lru.cache[key]
	if !ok {
		return "", false
	}

	// remove node from current position
	if (node.prev != nil) || (node.next != nil) {
		node.prev.next = node.next
		node.next.prev = node.prev
	}

	// put node to head
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node

	return node.value, true
}

func (lru *LRUCache) PrintCache() {
	current := lru.head.next
	for current != lru.tail {
		fmt.Printf("%s: %s \n", current.key, current.value)
		current = current.next
	}
	fmt.Printf("(size: %d/%d)\n", len(lru.cache), lru.capacity)
}
