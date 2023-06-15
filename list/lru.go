package list

import "container/list"

type Node struct {
	key, val int
	ele      *list.Element
}

type LRUCache struct {
	size int
	mp   map[int]*Node
	list *list.List
}

func NewLRUCache(size int) *LRUCache {
	l := &LRUCache{
		size: size,
		mp:   make(map[int]*Node),
		list: list.New(),
	}

	return l
}

func (l *LRUCache) Put(key, val int) {
	v, found := l.mp[key]
	if found {
		l.delete(v)
	}

	if l.size == len(l.mp) {
		l.deleteTail()
	}

	l.insert(&Node{key: key, val: val})
}

func (l *LRUCache) Get(key int) int {
	v, found := l.mp[key]
	if found {
		l.list.MoveToFront(v.ele)
		return v.val
	}

	return -1
}

func (l *LRUCache) insert(node *Node) {
	l.mp[node.key] = node
	node.ele = l.list.PushFront(node)
}

func (l *LRUCache) delete(node *Node) {
	delete(l.mp, node.key)
	l.list.Remove(node.ele)
}

func (l *LRUCache) deleteTail() {
	tail := l.list.Back()
	if tail != nil {
		l.delete(tail.Value.(*Node))
	}
}
