package cache

type LRUCache struct {
	nodeMap  map[int]*Node
	nodeList *List
	cap      int
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		nodeMap:  make(map[int]*Node, capacity),
		nodeList: NewList(),
		cap:      capacity,
	}
}

func (c *LRUCache) Get(key int) int {
	if _, ok := c.nodeMap[key]; !ok {
		return -1
	}

	c.nodeList.Remove(c.nodeMap[key])
	c.nodeList.PushFront(c.nodeMap[key])

	return c.nodeMap[key].Val
}

func (c *LRUCache) Put(key int, value int) {
	if curr, ok := c.nodeMap[key]; ok {
		delete(c.nodeMap, key)
		c.nodeList.Remove(curr)
		c.Put(key, value)
		return
	}

	node := &Node{
		prev: nil,
		next: nil,
		Key:  key,
		Val:  value,
	}

	if len(c.nodeMap) == c.cap {
		curr := c.nodeList.PopBack()
		delete(c.nodeMap, curr.Key)
	}

	c.nodeList.PushFront(node)
	c.nodeMap[key] = node
}
