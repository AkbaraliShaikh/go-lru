package main

import (
	"fmt"
	"go-lru/list"
)

func main() {
	fmt.Println("LRU using doubly linked list")
	lru := NewLRUCache(2)

	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))

	lru.Put(3, 3)
	fmt.Println(lru.Get(2))

	lru.Put(4, 4)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(4))

	fmt.Println("LRU using go inbuilt doubly linked list (container/list)")
	llru := list.NewLRUCache(2)

	llru.Put(1, 1)
	llru.Put(2, 2)
	fmt.Println(llru.Get(1))

	llru.Put(3, 3)
	fmt.Println(llru.Get(2))

	llru.Put(4, 4)
	fmt.Println(llru.Get(1))
	fmt.Println(llru.Get(3))
	fmt.Println(llru.Get(4))
}

// output
// LRU using doubly linked list
// 1
// -1
// -1
// 3
// 4

// LRU using go inbuilt doubly linked list (container/list)
// 1
// -1
// -1
// 3
// 4
