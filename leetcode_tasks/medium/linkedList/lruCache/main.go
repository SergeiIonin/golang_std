package main

import "fmt"

// https://leetcode.com/problems/lru-cache/

// Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

// Implement the LRUCache class:

// LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
// int get(int key) Return the value of the key if the key exists, otherwise return -1.
// void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache.
// If the number of keys exceeds the capacity from this operation, evict the least recently used key.
// The functions get and put must each run in O(1) average time complexity.

func main() {
	// generate test case for LRU cache with capacity 2 and perform operations on it
	// ["LRUCache","put","put","get","put","get","put","get","get","get"]
	// [[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
	// expected output: [null,null,null,1,null,-1,null,-1,3,4]

	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1)) // 1
	lru.Put(3, 3)
	fmt.Println(lru.Get(2)) // -1
	lru.Put(4, 4)
	fmt.Println(lru.Get(1)) // -1
	fmt.Println(lru.Get(3)) // 3
	fmt.Println(lru.Get(4)) // 4

	fmt.Println("------------------")

	lru1 := Constructor(1)
	lru1.Put(2, 1)
	fmt.Println(lru1.Get(2)) // 1
	lru1.Put(3, 2)
	fmt.Println(lru1.Get(2)) // -1
	fmt.Println(lru1.Get(3)) // 2

	fmt.Println("------------------")
	lru2 := Constructor(2)
	lru2.Put(2, 1)
	lru2.Put(1, 1)
	lru2.Put(2, 3)
	lru2.Put(4, 1)
	fmt.Println(lru2.Get(1)) // -1
	fmt.Println(lru2.Get(2)) // 3

	fmt.Println("------------------")
	lru3 := Constructor(3)
	lru3.Put(1, 1)
	lru3.Put(2, 2)
	lru3.Put(3, 3)

	lru3.Put(4, 4)

	fmt.Println(lru3.Get(4)) // 4
	fmt.Println(lru3.Get(3)) // 3
	fmt.Println(lru3.Get(2)) // 2
	fmt.Println(lru3.Get(1)) // -1
	lru3.Put(5, 5)
	fmt.Println(lru3.Get(1)) // -1
	fmt.Println(lru3.Get(2)) // 2
	fmt.Println(lru3.Get(3)) // 3
	fmt.Println(lru3.Get(4)) // -1
	fmt.Println(lru3.Get(5)) // 5

}

type LRUCache struct {
	maxCapacity int
	capacity    int
	head        *Node
	latest      *Node
	hashmap     map[int]*Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		maxCapacity: capacity,
		capacity:    0,
		head:        nil,
		latest:      nil,
		hashmap:     make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.hashmap[key]; ok {
		if node == this.latest {
			return node.Val
		}
		prev := node.Prev
		next := node.Next
		if prev != nil {
			prev.Next = next
			next.Prev = prev
			node.Next = nil
			this.latest.Next = node
			node.Prev = this.latest
			this.latest = node
		} else { // hence node == head
			this.head = next
			this.head.Prev = nil
			node.Prev = this.latest
			node.Next = nil
			this.latest.Next = node
			this.latest = node
		}

		return node.Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.hashmap[key]; ok {
		node.Val = value
		if node != this.latest {
			prev := node.Prev
			next := node.Next
			if prev != nil {
				prev.Next = next
				next.Prev = prev
				node.Next = nil
				node.Prev = this.latest
				this.latest.Next = node
				this.latest = node
			} else { // hence node == head
				this.head = next
				this.head.Prev = nil
				node.Prev = this.latest
				node.Next = nil
				this.latest.Next = node
				this.latest = node
			}
		}
	} else if this.head != nil {
		if this.capacity != this.maxCapacity {
			node = &Node{key, this.latest, nil, value}
			this.latest.Next = node
			this.latest = node
			this.hashmap[key] = node
			this.capacity++
		} else {
			if this.maxCapacity == 1 {
				delete(this.hashmap, this.head.Key)
				node = &Node{key, nil, nil, value}
				this.head = node
				this.latest = node
				this.hashmap[key] = node
			} else {
				tail := this.head.Next
				delete(this.hashmap, this.head.Key)
				this.head = tail
				this.head.Prev = nil
				node = &Node{key, this.latest, nil, value}
				this.latest.Next = node
				this.latest = node
				this.hashmap[key] = node
			}
		}
	} else {
		node = &Node{key, nil, nil, value}
		this.head = node
		this.latest = node
		this.hashmap[key] = node
		this.capacity++
	}
}

type Node struct {
	Key  int
	Prev *Node
	Next *Node
	Val  int
}

func printNode(node *Node) {
	var accum []int
	if node == nil {
		fmt.Println("[]")
	}
	for {
		accum = append(accum, node.Val)
		node = node.Next
		if node == nil {
			break
		}
	}
	fmt.Println(accum)
}
