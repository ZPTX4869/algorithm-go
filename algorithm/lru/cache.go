package lru

type LRUCache struct {
	m   map[int]*Node
	l   *List
	cap int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		m:   make(map[int]*Node, capacity),
		l:   NewList(),
		cap: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.m[key]; !ok {
		return -1
	}

	this.l.Remove(this.m[key])
	this.l.PushFront(this.m[key])

	return this.m[key].Val
}

func (this *LRUCache) Put(key int, value int) {
	if curr, ok := this.m[key]; ok {
		delete(this.m, key)
		this.l.Remove(curr)
		this.Put(key, value)
		return
	}

	node := &Node{
		prev: nil,
		next: nil,
		Key:  key,
		Val:  value,
	}

	if len(this.m) == this.cap {
		curr := this.l.PopBack()
		delete(this.m, curr.Key)
	}

	this.l.PushFront(node)
	this.m[key] = node
}
