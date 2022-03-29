package queues

import "log"

type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		inStack:  make([]int, 0),
		outStack: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.inStack) == 0 && len(this.outStack) == 0 {
		log.Fatalln("Queue is empty!")
	}

	if len(this.outStack) == 0 {
		for len(this.inStack) != 0 {
			size := len(this.inStack)
			this.outStack = append(this.outStack, this.inStack[size-1])
			this.inStack = this.inStack[:size-1]
		}
	}

	size := len(this.outStack)
	ret := this.outStack[size-1]
	this.outStack = this.outStack[:size-1]

	return ret
}

func (this *MyQueue) Peek() int {
	if len(this.inStack) == 0 && len(this.outStack) == 0 {
		log.Fatalln("Queue is empty!")
	}

	if len(this.outStack) == 0 {
		for len(this.inStack) != 0 {
			size := len(this.inStack)
			this.outStack = append(this.outStack, this.inStack[size-1])
			this.inStack = this.inStack[:size-1]
		}
	}

	size := len(this.outStack)
	ret := this.outStack[size-1]

	return ret
}

func (this *MyQueue) Empty() bool {
	if len(this.inStack) == 0 && len(this.outStack) == 0 {
		return true
	}

	return false
}
