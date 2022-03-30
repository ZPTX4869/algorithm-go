package lru

type Node struct {
	prev, next *Node
	Key, Val   int
}

func NewNode(key, val int) *Node {
	return &Node{
		Key: key,
		Val: val,
	}
}

type List struct {
	head *Node
	tail *Node
}

func NewList() *List {
	return &List{
		head: nil,
		tail: nil,
	}
}

func (l *List) PushFront(node *Node) {
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	node.next = l.head
	l.head.prev = node
	l.head = node
}

func (l *List) PushBack(node *Node) {
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	l.tail.next = node
	node.prev = l.tail
	l.tail = node
}

func (l *List) PopFront() *Node {
	temp := l.head

	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return temp
	}

	l.head = l.head.next
	l.head.prev = nil

	return temp
}

func (l *List) PopBack() *Node {
	temp := l.tail

	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return temp
	}

	l.tail = l.tail.prev
	l.tail.next = nil

	return temp
}

func (l *List) Remove(node *Node) {
	if node == l.head {
		if l.head != l.tail {
			l.head = l.head.next
			l.head.prev = nil
		} else {
			l.head, l.tail = nil, nil
		}
		return
	}

	if node == l.tail {
		l.PopBack()
		return
	}

	node.prev.next = node.next
	node.next.prev = node.prev
}
