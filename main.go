package main

import "fmt"

func main() {
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
}

// output
// 1
// -1
// -1
// 3
// 4
