package main

type Node struct {
	key, value int
	next, prev *Node
}

type LRUCache struct {
	cap        int
	mp         map[int]*Node
	head, tail *Node
}

func NewLRUCache(capacity int) *LRUCache {
	lru := &LRUCache{
		cap:  capacity,
		mp:   map[int]*Node{},
		head: &Node{},
		tail: &Node{},
	}

	lru.head.next = lru.tail
	lru.tail.prev = lru.head

	return lru
}

func (l *LRUCache) Get(key int) int {
	val := -1
	n, found := l.mp[key]
	if found {
		val = n.value
		l.delete(n)
		l.insert(n)
	}
	return val
}

func (l *LRUCache) Put(key, value int) {
	n, found := l.mp[key]
	if found {
		l.delete(n)
		l.insert(&Node{key: key, value: value})
	} else {
		if l.cap == len(l.mp) {
			l.delete(l.tail.prev)
		}
		l.insert(&Node{key: key, value: value})
	}
}

func (l *LRUCache) insert(node *Node) {
	l.mp[node.key] = node
	node.next = l.head.next
	node.next.prev = node
	node.prev = l.head
	l.head.next = node
}

func (l *LRUCache) delete(node *Node) {
	delete(l.mp, node.key)
	node.prev.next = node.next
	node.next.prev = node.prev
}
