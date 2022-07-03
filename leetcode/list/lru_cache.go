package list

import llist "container/list"

type info struct {
	key, val int
}

type LRUCache struct {
	nodeMap  map[int]*llist.Element
	nodeList *llist.List
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		nodeMap:  make(map[int]*llist.Element, 0),
		nodeList: llist.New(),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.nodeMap[key]; ok {
        this.nodeList.MoveToFront(node)
        return node.Value.(info).val
    }

    return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.nodeMap[key]; !ok {
		if len(this.nodeMap) == this.capacity {
            backNode := this.nodeList.Back()
			this.nodeList.Remove(backNode)
			delete(this.nodeMap, backNode.Value.(info).key)
		}

		node := this.nodeList.PushFront(info{key: key, val: value})
		this.nodeMap[key] = node
        return
	}

	this.nodeMap[key].Value = info{key: key, val: value}
    this.nodeList.MoveToFront(this.nodeMap[key])
}
