package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LruCache struct {
	data    map[int]*Node
	maxSize int
	size    int
	head    *Node
}

func (this *LruCache) Add(value int) *Node {
	headNode := this.head
	createdNode := Node{value: value, next: headNode, prev: nil}
	this.head = &createdNode

	if headNode != nil {
		headNode.prev = &createdNode
	}

	fmt.Printf("[CALLED] Add\n")
	fmt.Printf("Value: %d, Address: %p \n", value, &createdNode)
	return &createdNode
}

func (this *LruCache) Move(value int) bool {
	return true
}

func (this *LruCache) Remove() int {
	var lastNode *Node

	for lastNode = this.head; lastNode.next != nil; lastNode = lastNode.next {
		fmt.Printf("%d, %p\n", lastNode.value, lastNode)
		if lastNode.next.next == nil {
			value := lastNode.next.value
			fmt.Printf("last")
			lastNode.next = nil
			return value
		}
	}

	return -1
}

func (this *LruCache) changeHead(selectNode *Node) {
	fmt.Println("[CALLED] changeHead")
	selectNodeNext := selectNode.next
	selectNodePrev := selectNode.prev
	fmt.Printf("[changeHead()] Address: %p, Value: %d, PREV Address: %p, NEXT Address: %p\n", selectNode, selectNode.value, selectNodePrev, selectNodeNext)

	if selectNodeNext != nil {
		selectNodePrev.next = selectNodeNext
	} else {
		selectNodePrev.next = nil
	}

	if selectNodeNext != nil {
		selectNodeNext.prev = selectNodePrev
	}

	selectNode.next = this.head
	this.head.prev = selectNode

	this.head = selectNode

	selectNode.prev = nil
}

func (this *LruCache) Get(value int) int {
	_, exist := this.data[value]

	if !exist {
		createdNode := this.Add(value)
		this.data[value] = createdNode
		calcSize := this.size + 1

		if calcSize > this.maxSize {
			removedValue := this.Remove()
			delete(this.data, removedValue)
		} else {
			this.size += 1
		}
	} else {
		this.changeHead(this.data[value])
	}

	return value
}

func (this *LruCache) Show() {
	fmt.Println("[CALLED] SHOW()")
	fmt.Println("[Show()] maxSize: ", this.maxSize)
	fmt.Println("[Show()] size: ", this.size)
	fmt.Println("[Show()] data: ", this.data)
	msg := ""
	for node := this.head; node != nil; node = node.next {
		fmt.Printf("[SHOW()] Address: %p, Value: %d, PREV Address: %p, NEXT Address: %p\n", node, node.value, node.prev, node.next)
		msg += strconv.Itoa(node.value)
		if node.next != nil {
			msg += " -> "
		}
	}
	fmt.Println(msg)
	fmt.Printf("\n")

}

func LruCacheConstrucctor(maxSize int) LruCache {
	l := LruCache{data: nil, size: 0, maxSize: maxSize, head: nil}
	l.data = map[int]*Node{}
	return l
}

func case1() {
	fmt.Println("================= CASE 1 START=================")
	lruCache := LruCacheConstrucctor(10)
	fmt.Println("[CREATED] LRU Cached: ", lruCache)
	lruCache.Show()

	lruCache.Get(1)
	lruCache.Show()

	lruCache.Get(2)
	lruCache.Show()

	lruCache.Get(3)
	lruCache.Show()

	lruCache.Get(4)
	lruCache.Show()

	lruCache.Get(5)
	lruCache.Show()
	fmt.Println("================= CASE 1 END=================")
}
func case2() {
	fmt.Println("================= CASE 2 START=================")
	lruCache := LruCacheConstrucctor(10)
	fmt.Println("[CREATED] LRU Cached: ", lruCache)
	lruCache.Show()

	lruCache.Get(1)
	lruCache.Show()

	lruCache.Get(2)
	lruCache.Show()

	lruCache.Get(3)
	lruCache.Show()

	lruCache.Get(2)
	lruCache.Show()

	lruCache.Get(5)
	lruCache.Show()
	fmt.Println("================= CASE 2 END=================")
}

func case3() {
	fmt.Println("================= CASE 3 START=================")
	lruCache := LruCacheConstrucctor(2)
	fmt.Println("[CREATED] LRU Cached: ", lruCache)
	lruCache.Show()

	lruCache.Get(1)
	lruCache.Show()

	lruCache.Get(2)
	lruCache.Show()

	lruCache.Get(3)
	lruCache.Show()

	lruCache.Get(2)
	lruCache.Show()

	lruCache.Get(5)
	lruCache.Show()
	fmt.Println("================= CASE 3 END=================")
}
func case4() {
	fmt.Println("================= CASE 4 START=================")
	lruCache := LruCacheConstrucctor(1)
	fmt.Println("[CREATED] LRU Cached: ", lruCache)
	lruCache.Show()

	lruCache.Get(1)
	lruCache.Show()

	lruCache.Get(2)
	lruCache.Show()

	lruCache.Get(3)
	lruCache.Show()

	lruCache.Get(2)
	lruCache.Show()

	lruCache.Get(5)
	lruCache.Show()
	fmt.Println("================= CASE 4 END=================")
}

func main() {
	case1()
	case2()
	case3()
	case4()
}
