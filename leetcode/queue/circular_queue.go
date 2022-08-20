package queue

type MyCircularQueue struct {
	head int
	tail int
	nums []int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		head: 0,
		tail: 0,
		nums: make([]int, k+1),
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	this.nums[this.tail] = value
	this.tail = (this.tail + 1) % len(this.nums)

	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	this.head = (this.head + 1) % len(this.nums)

	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.nums[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.nums[(this.tail+len(this.nums)-1) % len(this.nums)]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == this.tail
}

func (this *MyCircularQueue) IsFull() bool {
	return (this.tail + 1) % len(this.nums) == this.head
}
