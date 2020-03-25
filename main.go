package main

import (
	"fmt"

	"./src/lru"
)

func case1() {
	fmt.Println("================= CASE 1 START=================")
	lruCache := lru.LruCacheConstrucctor(10)
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
	lruCache := lru.LruCacheConstrucctor(10)
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
	lruCache := lru.LruCacheConstrucctor(2)
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
	lruCache := lru.LruCacheConstrucctor(1)
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
