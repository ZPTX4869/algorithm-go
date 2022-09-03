package cache

import "container/list"

type LRU struct {
	len       int
	cap       int
	nodeList  *list.List
	keyToNode map[string]*list.Element
}

func NewLRU(cap int) LRU {
	return LRU{
		len:       0,
		cap:       cap,
		nodeList:  list.New(),
		keyToNode: make(map[string]*list.Element),
	}
}

func (lru *LRU) Put(key int, val int) bool {
	return false
}

func (lru *LRU) Get(key int) int {
	return 0
}
